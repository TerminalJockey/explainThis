package main

import (
	"io"
	"log"
	"net"

	_ "github.com/TerminalJockey/explainThis/one"
)

func main() {
	startWebServer()
}

func startWebServer() {
	log.Println("starting server")
	ln, err := net.Listen("tcp", "localhost:7777")
	if err != nil {
		log.Println(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
		}
		go handler(conn)
	}
}

func handler(conn net.Conn) {
	defer conn.Close()
	ret, err := io.ReadAll(conn)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(ret))

}
