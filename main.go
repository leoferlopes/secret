package main

import (
	"fmt"
	"os"
	"github.com/jessevdk/go-flags"
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

func _main() {
	var serverParams ServerParams
	var clientParams ClientParams

	args := flags.NewNamedParser("secret", flags.Options(flags.Default))
	args.AddCommand("client", "Run secret client", "Run secret client", &clientParams)
	args.AddCommand("server", "Run secret server", "Run secret client", &serverParams)
	_, err := args.ParseArgs(os.Args[1:])
	if err != nil {
		os.Exit(2)
	}

	switch args.Active.Name {
	case "server":
		server(serverParams)
	case "client":
		client(clientParams)
	}
}

func server(params ServerParams) {
	fmt.Printf("%+v\n", params)
}

func client(params ClientParams) {
	fmt.Printf("%+v\n", params)
}

func main()  {
	var message uint16 = 123

	publicKey := crypto.RSAKey{
		N: 2281,
		E: 29,
	}
	privateKey := crypto.RSAKey{
		N: 2281,
		E: 1625,
	}

	encrypted := crypto.RSACipher(message, privateKey)
	fmt.Println("encrypted",encrypted)
	decrypted := crypto.RSACipher(encrypted, publicKey)
	fmt.Println("decrypted", decrypted)
}