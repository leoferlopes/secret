package crypto

import (
	"errors"
	"fmt"
	"math/big"
)

func NewMACMessage(message CryptoMessage) CryptoMessage {
	return &MacMessage{
		CryptoMessage: message,
	}
}

type MacMessage struct {
	CryptoMessage
}

func (m *MacMessage) Encrypt() (CryptoMessage, error) {
	s, err := m.CryptoMessage.Encrypt()
	if err != nil {
		return nil, err
	}
	data := s.Bytes()
	println("---------------", len(data))
	hash := big.NewInt(int64(crc(data)))
	data = append(s.Bytes(), padding(hash.Bytes(), 4)...)
	data = append(data, padding(number(len(s.Bytes())).Bytes(), 4)...)
	fmt.Println("----------------", padding(number(len(s.Bytes())).Bytes(), 4))
	s.SetBytes(data)
	fmt.Println("++++++++++++++", len(data))
	return s, nil
}

func (m *MacMessage) Decrypt() (CryptoMessage, error) {
	s, err := m.CryptoMessage.Decrypt()
	if err != nil {
		return nil, err
	}
	bytes := s.Bytes()
	actualSize := int(number().SetBytes(bytes[len(bytes)-4:]).Int64())
	fmt.Println("-----------", actualSize)
	bytes = bytes[:len(bytes)-4]
	data := append([]byte{}, bytes[len(bytes)-actualSize-4:len(bytes)-4]...)
	println("---------------", len(data))
	hash := int64(crc(data))
	if hash == number().SetBytes(bytes[len(bytes)-4:]).Int64() {
		s.SetBytes(bytes[:len(bytes)-4])
		return s, nil
	} else {
		return nil, errors.New("Could not verify the message '" + s.String() + "'")
	}
}
