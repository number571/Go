package main 

import (
    "strings"
    "bufio"
    "net"
    "fmt"
    "os"
)

const (
    CMD_EXIT = ":exit"
    CMD_CONNECT = ":connect"
    CMD_DISCONNECT = ":disconnect"

    PROTOCOL_TCP = "tcp"
    PROTOCOL_UDP = "udp"

    DNS = "8.8.8.8:8080"
    BUFF = 1024
)

var (
    __main_port string = ""
    __nickname  string = "default"

    __local_ip string
    __to_ip_port []string

    __connected bool = false
    __size_connection uint16 = 0
)

func main () {
    check_args(os.Args)
    state_server()
}

func check_args (args []string) {
    var (
        nick bool = false
        port bool = false
    )

    for _, value := range args[1:] {

        if port { __main_port = value; port = false; continue }
        if nick { __nickname  = value; nick = false; continue }

        switch value {
            case "-p", "--port":
                port = true
            case "-n", "--nickname":
                nick = true
        }
    }

    if __main_port == "" { get_error("port not specified") }
}

func state_server () {
    var buffer []byte = make([]byte, BUFF)

    go to_client()

    for {
        listen, err := net.Listen(PROTOCOL_TCP, ":" + __main_port)
        check_error(err)

        conn, err := listen.Accept()
        if err != nil { 
            listen.Close()
            continue 
        }

        for {
            length, err := conn.Read(buffer)
            if err != nil || length == 0 { break }
            fmt.Print(string(buffer[:length]))
        }

        conn.Close()
        listen.Close()
    }
}

func to_client () {
    fmt.Println(get_local_ip())
    channel := make(chan bool)
    for {
        go state_client(channel)
        <- channel
    }
}

func state_client (channel chan bool) {
    var (
        message string
        message_without_space string
        splited_message []string 
    )

    message = input_string("")
    message_without_space = strings.Replace(message, " ", "", -1)
    splited_message = strings.Split(message_without_space, "->")

    switch splited_message[0] {

    case CMD_EXIT:
        os.Exit(0)

    case CMD_CONNECT:
        if len(splited_message) == 2 {
            __to_ip_port = append(__to_ip_port, splited_message[1])
            __size_connection++
            __connected = true
        }

    case CMD_DISCONNECT:
        if len(splited_message) == 1 {
            for i := uint16(0); i < __size_connection; i++ {
                __to_ip_port = remove(__to_ip_port, 0)
            }
            __size_connection = 0
            __connected = false

        } else {
            for i := uint16(0); i < __size_connection; i++ {
                if splited_message[1] == __to_ip_port[i] {
                    __to_ip_port = remove(__to_ip_port, i)
                    __size_connection--
                    break
                }
            }

            if __size_connection == 0 {
                __connected = false
            }
        }

    default:
        if  __connected  {
            for i := uint16(0); i < __size_connection; i++ {
                conn, err := net.Dial(PROTOCOL_TCP, __to_ip_port[i])
                if err != nil {
                    __to_ip_port = remove(__to_ip_port, i)
                    __size_connection--
                    if __size_connection == 0 {
                        __connected = false
                    }
                } else {
                    conn.Write([]byte(fmt.Sprintf("[%s/%s]: %s\n", __local_ip, __nickname, message)))
                    conn.Close()
                }
            }
        }
    }

    channel <- true
}

func get_local_ip () string {
    conn, err := net.Dial(PROTOCOL_UDP, DNS)
    check_error(err)

    __local_ip = strings.Split(conn.LocalAddr().String(), ":")[0]
    conn.Close()
    
    return "Your local IP adress: " + __local_ip
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

func remove(list []string, num uint16) []string {
    return append(list[:num], list[num+1:]...)
}

func check_error (err error) {
    if err != nil {
        get_error(err.Error())
    }
}

func get_error (err string) {
    fmt.Println("Error:", err)
    os.Exit(1)
}
