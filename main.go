package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    
    for {
        fmt.Print("Enter your name: ")
        name, err := reader.ReadString('\n')
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
            os.Exit(1)
        }

        name = strings.TrimSpace(name) 
        if name == "" {
            fmt.Println("Name cannot be empty. Please try again.")
            continue 
        }

        fmt.Printf("Hello, %s 👋!\n", name)
        break
    }
}
