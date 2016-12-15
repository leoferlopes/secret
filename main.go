// A concurrent prime sieve

package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/leoferlopes/secret/client"
	"github.com/leoferlopes/secret/server"
	"github.com/leoferlopes/secret/types"
	"os"
	"github.com/leoferlopes/secret/crypto"
)

func main() {
	var serverParams types.ServerParams
	var clientParams types.ClientParams
	var testParams types.TestParams

	args := flags.NewNamedParser("secret", flags.Options(flags.Default))
	args.AddCommand("client", "Run secret client", "Run secret client", &clientParams)
	args.AddCommand("server", "Run secret server", "Run secret client", &serverParams)
	args.AddCommand("test", "Run secret test", "Run secret test", &testParams)
	_, err := args.ParseArgs(os.Args[1:])
	if err != nil {
		os.Exit(2)
	}

	switch args.Active.Name {
	case "server":
		runServer(serverParams)
	case "client":
		runClient(clientParams)
	case "test":
		test()
	}
}

func runServer(params types.ServerParams) {
	secretServer := server.NewServer(params)
	fmt.Printf("Listening %s...\n", secretServer.Bind)
	secretServer.Run()
}

func runClient(params types.ClientParams) {
	secretClient := client.NewClient(params)
	fmt.Printf("Sending file %s to server %s\n", secretClient.File, secretClient.Server)
	secretClient.Run()
}

func _test() {
	m := crypto.NewMessage([]byte("banana"))
	fmt.Println(m.String())
	k := crypto.NewXOR()
	rsa := crypto.NewRSA(4)
	s := crypto.NewRSAMessage(crypto.NewXORMessage(crypto.NewMACMessage(m), k), rsa)
	ss, err := s.Encrypt()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(ss.Bytes())
	sss := crypto.NewMACMessage(crypto.NewXORMessage(crypto.NewRSAMessage(ss, rsa), k))
	sss, err = sss.Decrypt()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(sss.String())
}

func test() {
	cypher := crypto.NewStandartCypher()
	fmt.Printf("cypher: %+v\n", *cypher.RSA)
	bytes := []byte("myBanana")
	fmt.Println("string", string(bytes))
	fmt.Println("bytes", bytes)
	sequence := uint64(512)
	encrypted := cypher.Encrypt(bytes, sequence)
	fmt.Println("encrypted", encrypted)
	decrypted, seq := cypher.Decrypt(encrypted)
	fmt.Println("decrypted", decrypted)
	fmt.Println("seq", seq)
	fmt.Println("string", string(decrypted))
}
