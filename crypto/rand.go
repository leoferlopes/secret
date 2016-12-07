package crypto

import (
	"math/big"
	crypto "crypto/rand"
	"math/rand"
	"bytes"
	"time"
)

func init() {
	rand.Seed(int64(time.Now().UnixNano()))
}

func randomPrime() *big.Int {
	b := make([]byte, 16)
	crypto.Read(b)
	p, err := crypto.Prime(bytes.NewBuffer(b), 8)
	for err != nil {
		b = make([]byte, 16)
		crypto.Read(b)
		p, err = crypto.Prime(bytes.NewBuffer(b), 8)

	}
	return p
}


func number(n... int) *big.Int {
	return big.NewInt(int64(append(n, 0)[0]))
}