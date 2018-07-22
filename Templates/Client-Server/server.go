package main 

import (
	"os"
	"fmt"
	"net"
)

const (
	PROTOCOL = "tcp"
	PORT = ":8080"
)

func main () {
	listen, err := net.Listen(PROTOCOL, PORT)
	check_error(err)
	defer listen.Close()

	fmt.Println("Server is listening ... ")
	
	for {
		conn, err := listen.Accept()
		if err != nil { break }
		conn.Write([]byte("hello, world"))
		conn.Close()
	}
}

func check_error (err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
