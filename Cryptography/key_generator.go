package main

import (
    "fmt"
    "os"
    "strings"
    "crypto/rand"
)

func main () {
    fmt.Print("[C]reate|[R]ead: ")
    var choice string
    fmt.Scanf("%s", &choice)

    if (strings.ToUpper(choice) == "C") {
        create_key()
    } else if (strings.ToUpper(choice) == "R") {
        read_key()
    } else {
        get_error("mode is not found")
    }

    os.Exit(0)
}

func create_key () {
    fmt.Print("File name: ")
    var file_name string
    fmt.Scanf("%s", &file_name)

    fmt.Print("Max number: ")
    var max_num int32
    fmt.Scanf("%d", &max_num)

    if (max_num < 0) {
        get_error("Max number < 0")
    }

    save_key(file_name, generate_key(max_num))
}

func read_key () {
    fmt.Print("File name: ")
    var file_name string
    fmt.Scanf("%s", &file_name)

    var length int32 = get_length(file_name)

    var (
        position int32
        quantity int32
    )

    fmt.Println("Length of file:", length)
        
    fmt.Print("Position: ")
    fmt.Scanf("%d", &position)

    fmt.Print("Quantity: ")
    fmt.Scanf("%d", &quantity)

    if (length < position + quantity) {
        get_error("length of file < pos + quan")
    } else {
        fmt.Println(get_chars(
            file_name, 
            length,
            position, 
            quantity,
        ))
    }
}

func get_error (err string) {
    fmt.Println("Error:", err)
    os.Exit(1)
}

func check_error (err error) {
    if err != nil { 
        fmt.Println("Error:", err)
        os.Exit(1) 
    }
}

func get_chars (name string, length, pos, quan int32) string {
    file, err := os.Open(name)
    check_error(err)
    defer file.Close()

    var bytes []byte = make([]byte, length)
    _, err = file.Read(bytes)
    check_error(err)

    var read string
    var max int32 = pos + quan

    for i := pos; i < max; i++ {
        read += string(bytes[i])
    }

    return read
}

func get_length (name string) int32 {
    file, err := os.Open(name)
    check_error(err)
    defer file.Close()

    info, err := file.Stat()
    check_error(err)
    return int32(info.Size())
}

func save_key (name string, key []byte) {
    file, err := os.Create(name)
    check_error(err)
    defer file.Close()
    file.WriteString(string(key))
}

func generate_key (max int32) []byte {
    var slice []byte = make([]byte, max)
    _, err := rand.Read(slice)
    if err != nil { return nil }
    for max = max-1; max >= 0; max-- {
        slice[max] = slice[max] % 95 + 33
    }
    return slice
}
