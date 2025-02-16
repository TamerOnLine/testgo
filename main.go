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

        // قراءة المدخلات مع التعامل مع الأخطاء
        input, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Error reading input. Please try again.")
            continue
        }

        name := strings.TrimSpace(input) // إزالة الفراغات والمسافات غير الضرورية

        if name == "" {
            fmt.Println("Name cannot be empty. Please try again.")
            continue
        }

        fmt.Printf("Hello, %s 👋!\n", name)
        break
    }
}
