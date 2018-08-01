package main 

import (
    "time"
    "net"
    "fmt"
    "os"
)

const (
    PROTOCOL = "tcp"
    PORT = ":8080"
    BUFF = 1024
    QUAN = 5
)

var (
    connect [QUAN]net.Conn
    user uint8 = 0
)

func main () {
    listen, err := net.Listen(PROTOCOL, PORT)
    check_error(err)
    defer listen.Close()
    
    fmt.Println("Server is listening...")

    for {
        time.Sleep(time.Millisecond * 500)
        if user < QUAN {
            connect[user], err = listen.Accept()
            if err != nil { break }
            try_connection(connect[user])
            go connection(connect[user])
        }
    }

    fmt.Scanln()
}

func try_connection (conn net.Conn) {
    conn.Write([]byte("1"))
}

func connection (conn net.Conn) {
    user++
    var buffer []byte = make([]byte, BUFF)
    defer conn.Close()
    for {
        length, err := conn.Read(buffer)
        if err != nil { break }
        result := string(buffer[:length])
        for i := uint8(0); i < user; i++ {
            if connect[i] != conn {
                connect[i].Write([]byte(result))
            }
        }
    }
    user--
}

func check_error (err error) {
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
}
