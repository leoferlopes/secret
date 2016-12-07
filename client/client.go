package client

import (
	"github.com/leoferlopes/secret/types"
	"net"
	"os"
	"fmt"
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
	buf := make([]byte, 1024)

	// Read the incoming connection into the buffer.
	for length, err := file.Read(buf); length > 0; length, err = file.Read(buf) {
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}

		// Write the buffer into the file
		conn.Write(buf[0:length])
		check(err)
	}
	file.Close()
	conn.Close()
}