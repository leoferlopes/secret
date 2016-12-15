package crypto

import (
	"encoding/binary"
)

type StandartCypher struct {
	RSA *RSA
}

func (cypher *StandartCypher) Encrypt(b []byte, sequence uint64) []byte {
	s := make([]byte, 4)
	binary.LittleEndian.PutUint32(s, uint32(sequence))
	bytes := append(b, s...)
	m := NewMessage(bytes)
	m = NewMACMessage(m)
	m = NewRSAMessage(m, cypher.RSA)
	m, _ = m.Encrypt()
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
	sequence := binary.LittleEndian.Uint32(s)
	return bytes[:len(bytes)-4], uint64(sequence)
}

func NewStandartCypher() *StandartCypher {
	return &StandartCypher{
		RSA: NewRSAHardcoded(),
	}
}
