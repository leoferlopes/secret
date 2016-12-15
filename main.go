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

func test() {
	cypher := crypto.NewStandartCypher()
	bytes := []byte("Lorem ipsum dolor sit amet.")
	fmt.Println("text:", string(bytes))
	fmt.Println("bytes:", bytes)
	sequence := uint64(1)
	encrypted := cypher.Encrypt(bytes, sequence)
	fmt.Println("encrypted:", encrypted)
	decrypted, seq := cypher.Decrypt(encrypted)
	fmt.Println("decrypted:", decrypted)
	fmt.Println("sequence", seq)
	fmt.Println("text:", string(decrypted))
}
