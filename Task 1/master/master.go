package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	ip := "0.0.0.0:8080"
	// Start server
	listener, err := net.Listen("tcp", ip) // Listen for incoming connections on specified address and port
	if err != nil {
		fmt.Println("Error listening:", err.Error()) // Print error if unable to start listening
		return
	}
	defer listener.Close()                               // Close the listener when main function returns
	fmt.Println("Server started, listening on port", ip) // Print a message indicating server startup

	// Accept incoming connections
	for {
		connection, err := listener.Accept() // Accept incoming connection
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error()) // Print error if unable to accept connection
			continue
		}

		go handleConnection(connection) // Handle the connection in a separate goroutine
	}
}

func handleConnection(connection net.Conn) {
	defer connection.Close() // Close the connection when handleConnection function returns
	reader := bufio.NewReader(os.Stdin)

	deviceIP := connection.RemoteAddr().String()
	fmt.Println("New client connected:", deviceIP) // Print a message indicating a new client connection

	fmt.Print("Enter command (shutdown, restart): ")
	command, _ := reader.ReadString('\n')
	command = strings.TrimSpace(command)

	// Validate the command
	switch command {
	case "shutdown", "restart":
		// Send the provided command to the connected client
		_, err := connection.Write([]byte(command)) // Write the provided command to the connected client
		if err != nil {
			fmt.Println("Error sending command to device:", err.Error()) // Print error if unable to send command
			return
		}
	default:
		fmt.Println("Invalid command. Please enter shutdown, or restart.")
	}
}
