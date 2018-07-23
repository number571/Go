package main

import (
    "os"
    "fmt"
    "bufio"
    "strings"
    "time"
)

func main() {
    go start_timer(5)
    var message string = input_string("> ")
    fmt.Println(message)
}

func start_timer (num uint16) {
    time.Sleep(time.Second * time.Duration(num))
    fmt.Println("Time is out!")
    os.Exit(0)
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
