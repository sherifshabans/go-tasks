package main

import (
	"fmt"
	"net"
	"os/exec"
	"strings"
)

func main() {
	ip := "192.168.1.7:8080"
	// Connect to server
	connection, err := net.Dial("tcp", ip) // Connect to the server
	if err != nil {
		fmt.Println("Error connecting to server:", err.Error()) // Print error if unable to connect
		return
	}
	defer connection.Close() // Close the connection when main function returns

	fmt.Println("Connected to server")

	// Receive and execute commands continuously
	for {
		command := make([]byte, 1024)      // Create a buffer to store received command
		n, err := connection.Read(command) // Read command from the server
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Println("Error receiving command:", err.Error()) // Print error if unable to receive command
			}
			return
		}

		commandString := strings.TrimSpace(string(command[:n]))
		if commandString == "" {
			continue // Skip empty commands
		}

		// Execute command
		switch commandString {
		case "shutdown":
			err := exec.Command("shutdown", "/s", "/t", "0").Run() // Execute shutdown command
			if err != nil {
				fmt.Println("Failed to initiate shutdown:", err) // Print error if shutdown command execution fails
				return
			}
			fmt.Println("Shutdown command executed successfully") // Print success message if shutdown command executed successfully
		case "restart":
			err := exec.Command("shutdown", "/r", "/t", "0").Run() // Execute restart command
			if err != nil {
				fmt.Println("Failed to initiate restart:", err) // Print error if restart command execution fails
				return
			}
			fmt.Println("Restart command executed successfully") // Print success message if Sleep command executed successfully
		default:
			fmt.Println("Unknown command:", commandString) // Print error message for unknown command
		}
	}
}
