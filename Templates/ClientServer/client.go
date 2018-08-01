package main 

import (
	"os"
	"fmt"
	"net"
)

const (
	PROTOCOL = "tcp"
	IPV4 = "127.0.0.1"
	PORT = ":8080"
	BUFF = 64
)

func main () {
	conn, err := net.Dial(PROTOCOL, IPV4 + PORT)
	check_error(err)
	defer conn.Close()

	var buffer []byte = make([]byte, BUFF)
	var message string

	for {
		length, err := conn.Read(buffer)
		if length == 0 || err != nil { break }
		message += string(buffer)
	}

	fmt.Println(message)
}

func check_error(err error) {
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
