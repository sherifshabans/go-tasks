package main

import (
	"io"
	"log"
	"net"
	"os"
)

// handleConnection handles the file transfer for each client connection
func handleConnection(conn net.Conn, file *os.File) {
	defer conn.Close()

	// Create a buffer to read the file chunk by chunk
	buffer := make([]byte, 1024)
	for {
		// Read a chunk from the file
		bytesRead, err := file.Read(buffer)
		if err != nil {
			if err != io.EOF {
				log.Println("Error reading file:", err)
			}
			break
		}

		// Write the chunk to the connection
		_, err = conn.Write(buffer[:bytesRead])
		if err != nil {
			log.Println("Error writing to connection:", err)
			break
		}
	}
}

func main() {
	// Open the file to be shared
	file, err := os.Open("file.txt")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	defer file.Close()

	// Listen for incoming connections on port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Error listening:", err)
	}
	defer listener.Close()

	log.Println("Server started. Listening on port 8080...")

	for {
		// Accept incoming connection
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		log.Println("Client connected:", conn.RemoteAddr())

		// Handle connection in a separate goroutine
		go handleConnection(conn, file)
	}
}
