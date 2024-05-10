package main

import (
	"log"
	"net"
	"os"
)

func main() {
	// Connect to the server
	connection, err := net.Dial("tcp", "localhost:8080") // Connect to the server
	if err != nil {
		log.Fatal("Error connecting to server:", err) // Print error if unable to connect
	}
	defer connection.Close() // Close the connection when main function returns

	log.Println("Connected to server.") // Print message indicating successful connection

	// Create a file to write the received data
	file, err := os.Create("newFile.txt") // Create a new file to write received data
	if err != nil {
		log.Fatal("Error creating file:", err) // Print error if unable to create file
	}
	defer file.Close() // Close the file when main function returns

	// Specify the number of bytes to receive
	bytesToReceive := 4 // Example: Receive 1024 bytes

	// Read data from connection and write it to file
	buffer := make([]byte, bytesToReceive)        // Create a buffer to store received data
	bytesReceived, err := connection.Read(buffer) // Read data from the connection
	if err != nil {
		log.Fatal("Error receiving data:", err) // Print error if unable to receive data
	}

	// Write received bytes to file
	_, err = file.Write(buffer[:bytesReceived]) // Write received data to the file
	if err != nil {
		log.Fatal("Error writing to file:", err) // Print error if unable to write to file
	}

	log.Println("File received successfully.") // Print message indicating successful file reception
}
