package main 

import (
    "strings"
    "bufio"
    "net"
    "fmt"
    "os"
)

const (
    PROTOCOL = "tcp"
    IPV4 = "127.0.0.1"
    PORT = ":8080"
    BUFF = 1024
)

var (
    buffer []byte = make([]byte, BUFF)
    username string
)

func main () {
    conn, err := net.Dial(PROTOCOL, IPV4 + PORT)
    check_error(err)
    defer conn.Close()

    username = input_string("Nickname: ")
    
    send_request(conn)
    get_response(conn)
}

func send_request (conn net.Conn) {
    channel := make(chan bool)
    go func() {
        for {
            go send(conn, channel)
            <- channel
        }
    }()
}

func send (conn net.Conn, channel chan bool) {
    message := input_string("> ")
    if len(message) != 0 { 
        conn.Write([]byte(
            fmt.Sprintf("[%s]: %s\n> ", username, message),
        ))
    }
    channel <- true
}

func get_response (conn net.Conn) {
    channel := make(chan bool)
    for {
        go get(conn, channel)
        <- channel
    }
}

func get (conn net.Conn, channel chan bool) {
    length, _ := conn.Read(buffer)
    if length != 0 {
        fmt.Print(string(buffer[:length]))
    }
    channel <- true
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

func check_error (err error) {
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
}
