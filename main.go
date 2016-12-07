// A concurrent prime sieve

package main

import (
	"github.com/leoferlopes/secret/crypto"
	"fmt"
)

func main() {
	m := crypto.NewMessage([]byte("banana"))
	fmt.Println(m.String())
	k := crypto.Key([]byte("maca"))
	rsa := crypto.NewRSA(4)
	s := crypto.NewRSAMessage(crypto.NewXORMessage(crypto.NewMACMessage(m, k), k), rsa)
	ss, err := s.Encrypt()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(ss.Bytes())
	sss := crypto.NewMACMessage(crypto.NewXORMessage(crypto.NewRSAMessage(ss, rsa), k), k)
	sss, err = sss.Decrypt()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(sss.String())

}

