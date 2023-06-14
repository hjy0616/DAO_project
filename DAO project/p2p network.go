package main

import (
	"fmt"
	"log"
	"net"
)

const (
	connHost = "localhost"
	connPort = "8080"
	connType = "tcp"
)

func main() {
	// Listen for incoming connections.
	listener, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		log.Fatal("Error starting TCP server.")
	}
	defer listener.Close()

	// Listen for an incoming connection.
	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Error accepting connection.")
	}

	// Handle connections in a new goroutine.
	go handleConnection(conn)
}

func handleConnection(conn net.Conn) {
	// Read incoming data.
	buffer := make([]byte, 1024)
	length, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	// Handle the incoming data (i.e., a new block).
	handleBlock(buffer[:length])

	// Close the connection when you're done with it.
	conn.Close()
}

func handleBlock(blockData []byte) {
	// Here you would handle the new block. For now, we just print the data.
	fmt.Println(string(blockData))
}
