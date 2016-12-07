package crypto

import (
	"hash/crc32"
	"math/big"
	"fmt"
)

type OMessage []byte

func (m OMessage) String() string {
	return string(m)
}

func (m OMessage) xorEncrypt(key []byte) OMessage {
	size := len(m) + len(key) - 1;
	size -= (size % len(key));
	message := make([]byte, size - len(m))
	message = append(message, m...)
	for i := 0; i < len(message); {
		for j := 0; j < len(key); j++ {
			message[i] = message[i] ^ key[j]
			i++
		}
	}
	return OMessage(append(message, padding(number(len(m)).Bytes(), 8)...))
}

func (m OMessage) xorDecrypt(key []byte) OMessage {
	fmt.Println([]byte(m[len(m) - 8:]))
	actualSize := int(number().SetBytes(m[len(m) - 8:]).Int64())
	message := append([]byte{}, m[:len(m) - 8]...)
	fmt.Println([]byte(message))
	for i := 0; i < len(message); {
		for j := 0; j < len(key); j++ {
			message[i] = message[i] ^ key[j]
			i++
		}
	}
	return OMessage(message[len(message) - actualSize:])
}

func (m OMessage) macEncrypt(symmetricKey []byte) OMessage {
	data := append([]byte{}, m...)
	data = append(data, symmetricKey...)
	hash := big.NewInt(int64(crc32.ChecksumIEEE(data)))
	data = append(padding(hash.Bytes(), 4), m...)
	return OMessage(data)
}

func (m OMessage) macDecrypt(symmetricKey []byte) OMessage {
	data := append([]byte{}, m[4:]...)
	data = append(data, symmetricKey...)
	hash := int64(crc32.ChecksumIEEE(data))
	if hash == number(0).SetBytes(m[:4]).Int64() {
		return OMessage(m[4:])
	}
	panic("Could not verify the message" + m.String())
}

func (m OMessage) rsaEncrypt(publicKey []byte) OMessage {
	data := make([]byte, 0)
	for _, b := range m {
		c := number(0).Exp(number(0).SetBytes([]byte{b}), number(0).SetBytes(publicKey[2:]), number(0).SetBytes(publicKey[:2]))
		data = append(data, c.Bytes()...)
	}
	return OMessage(data)
}

func (m OMessage) rsaDecrypt(privateKey []byte) OMessage {
	data := make([]byte, 0)
	for index := 0; index < len(m); index += 2 {
		c := number(0).Exp(number(0).SetBytes(m[index:index + 2]), number(0).SetBytes(privateKey[2:]), number(0).SetBytes(privateKey[:2]))
		if index == len(m) - 2 {
			data = append(data, padding(c.Bytes(), 8)...)
		} else {
			data = append(data, c.Bytes()...)
		}
	}
	return OMessage(data)
}

func (m OMessage) Encrypt(symmetricKey []byte, publicKey []byte) OMessage {
	s := OMessage(m)
	//s = s.macEncrypt(symmetricKey)
	//s = s.xorEncrypt(symmetricKey)
	s = s.rsaEncrypt(publicKey)
	return s
}
func (m OMessage) Decrypt(symmetricKey []byte, secretKey []byte) OMessage {
	s := OMessage(m)
	s = m.rsaDecrypt(secretKey)
	//s = s.xorDecrypt(symmetricKey)
	//s = s.macDecrypt(symmetricKey)
	return s
}

type Message struct {
	message []byte
}

func (m *Message) Encrypt() (CryptoMessage, error) {
	return NewMessage(m.message), nil
}

func (m *Message) Decrypt() (CryptoMessage, error) {
	return NewMessage(m.message), nil
}

func (m *Message) String() string {
	return string(m.message)
}

func (m *Message) Bytes() []byte {
	return m.message
}

func (m *Message) SetBytes(message []byte) {
	m.message = message
}

func NewMessage(message []byte) CryptoMessage {
	return &Message{
		message: message,
	}
}

