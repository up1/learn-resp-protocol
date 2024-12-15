package main

import (
	"fmt"
	"net"
)

// Address and port to bind to
const addr = "0.0.0.0:6379"

func main() {
	// Binding to TCP port 6379
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	fmt.Println("TCP Server starting at ", addr)

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err)
			break
		}
		defer conn.Close()

		fmt.Println("Connection accepted from ", conn.RemoteAddr())
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	buf := make([]byte, 1024)
	for {
		// Read the incoming connection into the buffer
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading: ", err)
			break
		}

		if n == 0 {
			break
		}

		// Print the received data
		fmt.Println("Received: ", string(buf[:n]))
		conn.Write([]byte("+OK\r\n"))
	}
}
