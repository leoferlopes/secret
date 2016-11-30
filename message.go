package main

import (
	"hash/crc32"
	"math/big"
	"fmt"
)

type Message []byte

func (m Message) String() string {
	return string(m)
}

func (m Message) xor(key []byte) Message {
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
	return Message(message)
}

func (m Message) macEncrypt(symmetricKey []byte) Message {
	data := append([]byte{}, m...)
	data = append(data, symmetricKey...)
	hash := big.NewInt(int64(crc32.ChecksumIEEE(data)))
	data = append(padding(hash.Bytes(), 4),m...)
	return Message(data)
}

func (m Message) macDecrypt(symmetricKey []byte) Message {
	data := append([]byte{}, m[4:]...)
	data = append(data, symmetricKey...)
	hash := int64(crc32.ChecksumIEEE(data))
	fmt.Println([]byte(m[:4]))
	if hash == number(0).SetBytes(m[:4]).Int64() {
		return Message(m[4:])
	}
	println(hash, number(0).SetBytes(m[:4]).Int64())
	panic("dsds")
}

func (m Message) rsaEncrypt(publicKey []byte) Message {
	data := make([]byte, 0)
	for _, b := range m {
		c := number(0).Exp(number(0).SetBytes([]byte{b}), number(0).SetBytes(publicKey[2:]), number(0).SetBytes(publicKey[:2]))
		data = append(data, c.Bytes()...)
	}
	return Message(data)
}

func (m Message) rsaDecrypt(privateKey []byte) Message {
	data := make([]byte, 0)
	for index := 0; index < len(m); index += 2 {
		c := number(0).Exp(number(0).SetBytes(m[index:index + 2]), number(0).SetBytes(privateKey[2:]), number(0).SetBytes(privateKey[:2]))
		data = append(data, c.Bytes()...)
	}

	return Message(data)
}

func (m Message) Encrypt(symmetricKey []byte, publicKey []byte) Message {
	return m.macEncrypt(symmetricKey).xor(symmetricKey).rsaEncrypt(publicKey)
}
func (m Message) Decrypt(symmetricKey []byte, secretKey []byte) Message {
	return m.rsaDecrypt(secretKey).xor(symmetricKey).macDecrypt(symmetricKey)
}
