package crypto

import (
	"encoding/binary"
	"fmt"
)

type StandartCypher struct {
	RSA *RSA
}

func (cypher *StandartCypher) Encrypt(b []byte, sequence uint64) []byte {
	s := make([]byte, 4)
	binary.LittleEndian.PutUint32(s, uint32(sequence))
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
	s := bytes[len(bytes)-4:]
	fmt.Println("Decrypt.s", s)
	sequence := binary.LittleEndian.Uint32(s)
	fmt.Println("Decrypt.sequence", sequence)
	return bytes[:len(bytes)-4], uint64(sequence)
}

func NewStandartCypher() *StandartCypher {
	return &StandartCypher{
		RSA: NewRSAHardcoded(),
	}
}
