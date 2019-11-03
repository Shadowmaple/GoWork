package main

import (
	"io"
	"log"
	"net"
	"time"
)

func handleConnect(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().String()+"\n")
		if err != nil {
			log.Print(err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Server start")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
		}
		go handleConnect(conn)
	}
}
