package crypto

import (
	"encoding/binary"
	"fmt"
)

type StandartCypher struct {
	RSA *RSA
}

func (cypher *StandartCypher) Encrypt(b []byte, sequence uint64) []byte {
	s := make([]byte, 8)
	binary.LittleEndian.PutUint64(s, sequence)
	fmt.Println("Encrypt.s", s)
	bytes := append(b, s...)
	m := NewMessage(bytes)
	m = NewMACMessage(m)
	m = NewRSAMessage(m, cypher.RSA)
	m, _ = m.Encrypt()
	fmt.Println("Encrypt.m.Bytes()", m.Bytes())
	return m.Bytes()
}

func (cypher *StandartCypher) Decrypt(b []byte) ([]byte, uint64) {
	m := NewMessage(b)
	m = NewRSAMessage(m, cypher.RSA)
	m = NewMACMessage(m)
	m, err := m.Decrypt()
	if err != nil {
		panic(err)
	}
	bytes := m.Bytes()
	s := bytes[len(bytes) - 8:]
	fmt.Println("Decrypt.s", s)
	sequence := binary.LittleEndian.Uint64(s)
	fmt.Println("Decrypt.sequence", sequence)
	return bytes[:len(bytes) - 8], sequence
}

func NewStandartCypher() *StandartCypher {
	return &StandartCypher{
		RSA: NewRSAHardcoded(),
	}
}