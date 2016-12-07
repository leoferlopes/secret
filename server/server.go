package server

import (
	"net"
	"fmt"
	"os"
	"github.com/leoferlopes/secret/types"
	"github.com/leoferlopes/secret/util"
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
	// Open file to append
	file, err := os.Create(server.File)
	defer file.Close()

	check(err)

	// Transfer bytes from conn to file
	util.TransferBuffer(conn, file)

	// Close file and connection
	file.Close()
	conn.Close()
}
