package client

import (
	"bytes"
	"fmt"
	"github.com/leoferlopes/secret/crypto"
	"github.com/leoferlopes/secret/types"
	"github.com/leoferlopes/secret/util"
	"net"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Client struct {
	File   string
	Server string
}

func NewClient(params types.ClientParams) *Client {
	return &Client{
		File:   params.File,
		Server: params.Server,
	}
}

func (client *Client) Run() {
	conn, err := net.Dial("tcp", client.Server)
	check(err)

	defer conn.Close()
	client.handleConnection(conn)
}

func (client *Client) handleConnection(conn net.Conn) {
	// read in input from stdin
	file, err := os.Open(client.File)
	check(err)
	rsa := crypto.NewRSA(16)
	conn.Write(encrypt(util.BEGIN, rsa))
	fmt.Println(encrypt(util.BEGIN, rsa))
	nounce := decrypt(read(conn), rsa)
	x := nounce[0]
	nounce[0] = nounce[1]
	nounce[1] = x
	conn.Write(encrypt(nounce, rsa))
	buf := bytes.NewBufferString("")
	buf.ReadFrom(file)
	buf = bytes.NewBuffer(encrypt(buf.Bytes(), rsa))
	util.TransferBuffer(buf, conn)
	file.Close()
	conn.Close()
}

func encrypt(b []byte, rsa *crypto.RSA) []byte {
	m := crypto.NewMessage(b)
	m = crypto.NewMACMessage(m)
	m = crypto.NewRSAMessage(m, rsa)
	m, _ = m.Encrypt()
	return m.Bytes()
}

func decrypt(b []byte, rsa *crypto.RSA) []byte {
	m := crypto.NewMessage(b)
	m = crypto.NewRSAMessage(m, rsa)
	m = crypto.NewMACMessage(m)
	m, _ = m.Decrypt()
	return m.Bytes()
}

func read(conn net.Conn) []byte {
	buf := bytes.NewBufferString("")
	util.TransferBuffer(conn, buf)
	return buf.Bytes()
}
