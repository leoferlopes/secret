package client

import (
	"github.com/leoferlopes/secret/types"
	"net"
	"os"
	"github.com/leoferlopes/secret/util"
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
		File: params.File,
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

	util.TransferBuffer(file, conn)

	file.Close()
	conn.Close()
}