// A concurrent prime sieve

package main

import (
	"fmt"
	"github.com/leoferlopes/secret/crypto"
)


func main() {
	symmetricKey := []byte("123")
	rsa := crypto.NewRSA()
	m := crypto.Message([]byte("banana"))
	fmt.Println(m.String())
	s := m.Encrypt(symmetricKey, rsa.PublicKey)
	t := s.Decrypt(symmetricKey, rsa.SecretKey)
	fmt.Println(t.String())
}

