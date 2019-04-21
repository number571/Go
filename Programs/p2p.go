package main 

import (
    "os"
    "fmt"
    "net"
    "bufio"
    "strings"
    "encoding/json"
)

const (
    PROTOCOL = "tcp"
    BUFF = 256
)

type packageTCP struct {
    From string
    Body string
}

var (
    connections = make(map[string]bool)
    address string
)

func main() {
    initArgs(os.Args)
    go client()
    server()
}

func initArgs(args []string) {
    var flag_address bool
    for _, value := range args {
        switch value {
            case "-a", "--address": 
                flag_address = true
                continue
        }
        switch {
            case flag_address: 
                address = value
                flag_address = false
        }
    }
    if address == "" { 
        printError("address undefined")
    }
}

func client() {
    for {
        var message = inputString()
        var splited = strings.Split(message, " ")
        switch splited[0] {
            case ":exit": os.Exit(0)
            case ":network": network()
            case ":connect":
                if len(splited) > 1 {
                    connectTo(splited[1:])
                }
            case ":disconnect":
                if len(splited) > 1 {
                    disconnectFrom(splited[1:])
                }
            default: 
                for addr := range connections {
                    sendPacket(addr, message)
                }
        }
    }
}

func server() {
    listen, err := net.Listen(PROTOCOL, address)
    if err != nil {
        printError("can't run listener")
    }
    defer listen.Close()
    fmt.Println("[ClientP2P is run]")
    for {
        conn, err := listen.Accept()
        if err != nil {
            break
        }
        go handleConnect(conn)
    }
}

func handleConnect(conn net.Conn) {
    defer conn.Close()
    var (
        buffer = make([]byte, BUFF)
        message string
        pack packageTCP
    )
    for {
        length, err := conn.Read(buffer)
        if err != nil { break }
        message += string(buffer[:length])
    }
    json.Unmarshal([]byte(message), &pack)
    connectTo([]string{pack.From})
    fmt.Printf("[%s]: %s\n", pack.From, pack.Body)
}

func sendPacket(to, message string) {
    conn, err := net.Dial(PROTOCOL, to)
    if err != nil {
        disconnectFrom([]string{to})
        return
    }
    var pack = packageTCP {
        From: address,
        Body: message,
    }

    data, err := json.Marshal(pack)
    if err != nil {
        printError("can't convert pack to json")
    }

    conn.Write(data)
    conn.Close()
}

func network() {
    for addr := range connections {
        fmt.Println("|", addr)
    }
}

func connectTo(conn []string) {
    for _, value := range conn {
        connections[value] = true
    }
}

func disconnectFrom(conn []string) {
    for _, value := range conn {
        delete(connections, value)
    }
}

func inputString() string {
    message, _ := bufio.NewReader(os.Stdin).ReadString('\n')
    return strings.Replace(message, "\n", "", -1)
}

func printError(err string) {
    fmt.Println("[Error]:", err)
}
