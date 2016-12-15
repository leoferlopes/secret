package server

import (
	"fmt"
	"github.com/leoferlopes/secret/crypto"
	"github.com/leoferlopes/secret/types"
	"github.com/leoferlopes/secret/util"
	"net"
	"os"
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
	file, err := os.Create(server.File)
	check(err)
	cypher := crypto.NewStandartCypher()
	util.TransferDecryptBuffer(conn, file, cypher)
	file.Close()
	conn.Close()
}
