package main

import (
    "strings"
    "bufio"
    "fmt"
    "net"
    "os"  
)

const (
    PROTOCOL = "tcp"
    IPV4 = "127.0.0.1"
    PORT = ":8080"
    BUFF = 2048
)

func main() {

    var (
        conn net.Conn
        command string
        buff []byte
        length int
        err error
    )

    conn, err = net.Dial(PROTOCOL, IPV4 + PORT) 
    check_error(err)
    defer conn.Close()

    buff = make([]byte, BUFF)

    for {
        command = input_string("$ ")

        length, err = conn.Write([]byte(command))
        if length == 0 || err != nil { continue }

        length, err = conn.Read(buff)
        if length == 0 || err != nil { break }

        fmt.Println(string(buff[:length]))
    }

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

func check_error (err error) {
    if err != nil {
        fmt.Println("Error:",err) 
        os.Exit(1)
    } 
}
