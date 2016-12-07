// A concurrent prime sieve

package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
	"github.com/leoferlopes/secret/crypto"
)

type ClientParams struct {
	File   string `short:"f" long:"file" description:"File to transfer to server" value-name:"FILE" required:"true"`
	Server string `short:"s" long:"server" description:"Server address" value-name:"SERVER" required:"true"`
}

type ServerParams struct {
	File string `short:"f" long:"file" description:"File to transfer to server" value-name:"FILE" required:"true"`
	Port int    `short:"p" long:"port" description:"Port to listen" value-name:"PORT" required:"true"`
}

type TestParams struct {
}

func main() {
	var serverParams ServerParams
	var clientParams ClientParams
	var testParams TestParams

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
		server(serverParams)
	case "client":
		client(clientParams)
	case "test":
		test()
	}
}

func server(params ServerParams) {
	fmt.Printf("%+v\n", params)
}

func client(params ClientParams) {
	fmt.Printf("%+v\n", params)
}

func test() {
	symmetricKey := []byte("123")
	rsa := crypto.NewRSA()
	m := crypto.Message([]byte("banana"))
	fmt.Println(m.String())
	s := m.Encrypt(symmetricKey, rsa.PublicKey)
	t := s.Decrypt(symmetricKey, rsa.SecretKey)
	fmt.Println(t.String())
}