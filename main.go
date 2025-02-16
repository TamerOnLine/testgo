package main

import (
	"bufio"  // Package for buffered input handling
	"fmt"    // Package for formatted I/O
	"os"     // Package for interacting with the operating system
	"strings" // Package for string manipulation
)

func main() {
	reader := bufio.NewReader(os.Stdin) // Creates a buffered reader to read user input

	for {
		fmt.Print("Enter your name: ") // Displays a prompt to the user

		input, err := reader.ReadString('\n') // Reads input from the user until a newline character is found
		if err != nil { // Checks if an error occurred during input reading
			fmt.Println("Error reading input. Please try again.") // Displays an error message
			continue // Continues the loop to ask for input again
		}

		name := strings.TrimSpace(input) // Removes leading and trailing spaces or newline characters

		if name == "" { // Checks if the input is empty after trimming
			fmt.Println("Name cannot be empty. Please try again.") // Displays an error message if input is empty
			continue // Continues the loop to ask for input again
		}

		fmt.Printf("Hello, %s 👋!\n", name) // Prints a greeting message including the user's name
		break // Exits the loop since the input is valid
	}
}
