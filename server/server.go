package server

import (
	"net"
	"fmt"
	"os"
	"github.com/leoferlopes/secret/types"
)

type Server struct {
	File string
	Bind string
}

func NewServer(params types.ServerParams) *Server {
	return &Server{
		File: params.File,
		Bind: fmt.Sprintf(":%d", params.Port),
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (server *Server) Run() {
	ln, err := net.Listen("tcp", server.Bind)
	check(err)

	conn, err := ln.Accept()
	check(err)

	defer conn.Close()
	server.handleConnection(conn)
}

func (server *Server) handleConnection(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 1024)

	// Open file to append
	file, err := os.Create(server.File)
	defer file.Close()

	check(err)

	// Read the incoming connection into the buffer.
	for length, err := conn.Read(buf); length > 0; length, err = conn.Read(buf) {
		if err != nil {
			fmt.Println("Error reading:", err.Error())
		}

		// Write the buffer into the file
		file.Write(buf[0:length])
		check(err)
	}
	file.Close()

	// Send a response back to person contacting us.
	conn.Write([]byte("Message received.\n"))
	// Close the connection when you're done with it.
	conn.Close()
}
