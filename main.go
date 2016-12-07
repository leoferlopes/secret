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