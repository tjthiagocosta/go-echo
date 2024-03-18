package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// Listen for incoming connections
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Teardown the server when the program exits
	defer listener.Close()

	fmt.Println("Server is listening on port 8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		// Handle client connection in a new goroutine
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	// Close the connection after we're done
	defer conn.Close()

	// Read data
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		return
	}

	log.Println("Received data", buf[:n])

	// Write the same data back
	conn.Write(buf[:n])
}
