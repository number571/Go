package main

import (
    "os"
    "fmt"
    "bufio"
    "strings"
)

func main() {
    var message string = input_string("> ")
    fmt.Println(message)
}

func input_string (text string) string {
    var (
        reader *bufio.Reader
        command string
    )

    fmt.Print(text)

    reader = bufio.NewReader(os.Stdin)
    command, _ = reader.ReadString('\n')

    return strings.Replace(command, "\n", "", -1)
}
