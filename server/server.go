package server

import (
	"bytes"
	"fmt"
	"github.com/leoferlopes/secret/crypto"
	"github.com/leoferlopes/secret/types"
	"github.com/leoferlopes/secret/util"
	"math/rand"
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
	rsa := crypto.NewRSA(16)

	begin := decrypt(read(conn), rsa)
	println("hello")
	if begin[0] == util.BEGIN[0] {
		nounce := make([]byte, 2)
		rand.Read(nounce)
		conn.Write(encrypt(nounce, rsa))
		nounceResult := decrypt(read(conn), rsa)
		if nounce[0] == nounceResult[1] && nounce[1] == nounceResult[0] {
			file, err := os.Create(server.File)
			check(err)

			defer file.Close()
			// Transfer bytes from conn to file
			buf := bytes.NewBufferString("")
			util.TransferBuffer(conn, buf)
			file.Write(decrypt(buf.Bytes(), rsa))
		}
	}
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
	x, _ := os.Create("x")
	buf := bytes.NewBufferString("")
	println("a")
	util.TransferBuffer(conn, x)
	println("a")
	buf.ReadFrom(x)
	return buf.Bytes()
}
