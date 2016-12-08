package crypto

import (
	"errors"
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
	hash := big.NewInt(int64(crc(data)))
	data = append(s.Bytes(), padding(hash.Bytes(), 8)...)
	s.SetBytes(data)
	return s, nil
}

func (m *MacMessage) Decrypt() (CryptoMessage, error) {
	s, err := m.CryptoMessage.Decrypt()
	if err != nil {
		return nil, err
	}
	bytes := s.Bytes()
	data := append([]byte{}, bytes[:len(bytes)-8]...)
	hash := int64(crc(data))
	if hash == number().SetBytes(bytes[len(bytes)-8:]).Int64() {
		s.SetBytes(bytes[:len(bytes)-8])
		return s, nil
	} else {
		return nil, errors.New("Could not verify the message '" + s.String() + "'")
	}
}
