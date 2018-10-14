package main 

/*
    ./main -p 8080 -n user
        |:connect -> 127.0.0.1:9090
        |hello, world
        |:network
        |:file -> main.go
        |:disconnect
        |:exit
*/

import (
    "os/exec"
    "strings"
    "bufio"
    "net"
    "fmt"
    "os"
)

const (
    CMD_LINE        = "->"
    CMD_EXIT        = ":exit"
    CMD_FILE        = ":file"
    CMD_NETWORK     = ":network"
    CMD_CONNECT     = ":connect"
    CMD_DISCONNECT  = ":disconnect"

    CHECK_FILE      = "[:~:-file-:~:]"
    CHECK_NAME      = "[:~:-name-:~:]"

    PROTOCOL_TCP = "tcp"
    PROTOCOL_UDP = "udp"

    DNS  = "8.8.8.8:8080"
    BUFF = 512
)

var (
    __main_port string = ""
    __nickname  string = "..."

    __local_ip string = get_local_ip()
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
    var (
        buffer  []byte = make([]byte, BUFF)
        splited []string
        content string = ""
    )

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
            content += string(buffer[:length])
        }

        if strings.HasPrefix(content, CHECK_FILE) {
            splited = strings.Split(strings.Replace(content, CHECK_FILE, "", 1), CHECK_NAME)
            if len(splited) == 2 {
                input_to_file(splited[0], splited[1])
                fmt.Printf("[File '%s' saved]\n", splited[0])
            }
        } else {
            fmt.Print(content)
        }

        content = ""

        conn.Close()
        listen.Close()
    }
}

func to_client () {
    fmt.Println("Your local IP adress: " + __local_ip)
    for {
        state_client()
    }
}

func state_client () {
    var (
        message string
        splited []string 
    )

    message = input_string("")
    splited = strings.Split(strings.Replace(message, " ", "", -1), CMD_LINE)

    switch splited[0] {

    case CMD_EXIT:
        os.Exit(0)

    case CMD_FILE:
        if len(splited) == 2 {
            var content string = get_file_content(splited[1])
            if __connected {
                for index, value := range __to_ip_port {
                    conn, err := net.Dial(PROTOCOL_TCP, value)
                    if err != nil {
                        __to_ip_port = remove(__to_ip_port, index)
                        __size_connection--
                        if __size_connection == 0 {
                            __connected = false
                        }
                    } else {
                        conn.Write([]byte(CHECK_FILE + splited[1] + CHECK_NAME + content))
                        conn.Close()
                    }
                }
            }
        }

    case CMD_NETWORK:
        fmt.Println("==================")
        fmt.Println("Connection:", __connected)
        for _, value := range __to_ip_port {
            fmt.Println(value)
        }
        fmt.Println("==================")

    case CMD_CONNECT:
        if len(splited) == 2 {

            var flag bool = false

            for _, value := range __to_ip_port {
                if value == splited[1] {
                    flag = true
                    break
                }
            }

            if !flag {
                conn, err := net.Dial(PROTOCOL_TCP, splited[1])
                if err == nil {
                    __to_ip_port = append(__to_ip_port, splited[1])
                    __size_connection++
                    __connected = true

                    conn.Write([]byte(fmt.Sprintf("[User '%s:%s/%s' join]\n", __local_ip, __main_port, __nickname)))
                    conn.Close()
                }
            }
        }

    case CMD_DISCONNECT:
        if len(splited) == 1 {
            for __size_connection != 0 {
                __to_ip_port = remove(__to_ip_port, 0)
                __size_connection--
            }
            __connected = false

        } else {
            for index, value := range __to_ip_port {
                if splited[1] == value {
                    __to_ip_port = remove(__to_ip_port, index)
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
            for index, value := range __to_ip_port {
                conn, err := net.Dial(PROTOCOL_TCP, value)
                if err != nil {
                    __to_ip_port = remove(__to_ip_port, index)
                    __size_connection--
                    if __size_connection == 0 {
                        __connected = false
                    }
                } else {
                    conn.Write([]byte(fmt.Sprintf("[%s:%s/%s]: %s\n", __local_ip, __main_port, __nickname, message)))
                    conn.Close()
                }
            }
        }
    }
}

func system (line_command string) ([]byte, error) {
    var command []string = strings.Split(line_command, " ")
    return exec.Command(command[0], command[1:]...).Output()
}

func get_file_content (filename string) string {
    file, err := os.Open(filename)
    check_error(err)
    defer file.Close()

    var buffer []byte = make([]byte, 512)
    var message string = ""
    
    for {
        length, err := file.Read(buffer)
        if err != nil || length == 0 { break }
        message += string(buffer[:length])
    }

    return message
}

func get_local_ip () string {
    conn, err := net.Dial(PROTOCOL_UDP, DNS)
    check_error(err)

    var ip string = strings.Split(conn.LocalAddr().String(), ":")[0]
    conn.Close()
    
    return ip
}

func input_to_file (name, content string) error {
    file, err := os.Create(name)

    if err != nil {
        return err
    }

    file.Write([]byte(content))
    file.Close()

    return nil
}

func input_string (text string) string {
    var command string
    fmt.Print(text)
    
    command, _ = bufio.NewReader(os.Stdin).ReadString('\n')
    return strings.Replace(command, "\n", "", -1)
}

func remove(list []string, num int) []string {
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
