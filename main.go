// A concurrent prime sieve

package main

import (
	"crypto/rand"
	math_rand "math/rand"
	"bytes"
	"time"
	"fmt"
	"math/big"
)

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func Generate(ch chan <- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan <- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i % prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

// The prime sieve: Daisy-chain Filter processes.
func randomPrime() *big.Int {
	b := make([]byte, 16)
	rand.Read(b)
	p, err := rand.Prime(bytes.NewBuffer(b), 8)
	for err != nil {
		b = make([]byte, 16)
		rand.Read(b)
		p, err = rand.Prime(bytes.NewBuffer(b), 8)

	}
	return p
}

func number(b int64) *big.Int {
	return big.NewInt(b)
}

func main() {
	math_rand.Seed(int64(time.Now().UnixNano()))

	p := randomPrime()
	q := randomPrime()

	for number(0).Sub(p, q).Uint64() == 0 {
		q = randomPrime()
	}
	n := number(0).Mul(q, p)
	z := number(0).Mul(number(0).Sub(p, number(1)), number(0).Sub(q, number(1)))
	e := randomPrime()
	for number(0).Sub(e, n).Sign() != -1 {
		e = randomPrime()
	}
	d := randomPrime()
	for true {
		ed := big.NewInt(0)
		ed = number(0).Mul(e, d)
		m := number(0).Mod(ed, z)
		if m.Uint64() == 1 {
			break
		}
		d = d.Add(d, big.NewInt(1))
	}
	nb := padding(n.Bytes(), 2)
	db := padding(d.Bytes(), 2)
	eb := padding(e.Bytes(), 2)

	symmetricKey := []byte("such")
	publicKey := append(nb, eb...)
	secretKey := append(padding(n.Bytes(), 2), db...)
	m0 := []byte("baanananaan")
	size := len(m0) + len(symmetricKey) - 1;
	size -= (size % len(symmetricKey));
	m := Message(padding(m0, size))
	fmt.Println(m.String())
	s := m.Encrypt(symmetricKey, publicKey)
	t := s.Decrypt(symmetricKey, secretKey)
	fmt.Println(t.String())
}

func padding(array []byte, size int) []byte {
	pad := append(make([]byte, size), array...)
	return pad[len(pad) - size:]
}
