package client

import (
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
	file, err := os.Open(client.File)
	check(err)
	cypher := crypto.NewStandartCypher()
	util.TransferEncryptBuffer(file, conn, cypher)
	file.Close()
	conn.Close()
}
