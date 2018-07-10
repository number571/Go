package main

 /*
  * [Example]:
  *     $ go build main.go 
  *     [Generate]:
  *         $ ./main -g key.txt 1000
  *     [Read_Key]:
  *         $ ./main -r key.txt 50 100
  *     [Get_Length]:
  *         $ ./main -l key.txt 
  */

import (
    "fmt"
    "os"
    "strings"
    "strconv"
    "crypto/rand"
)

var modeG [2]string = [2]string{"-G", "--GENERATE"}
var modeR [2]string = [2]string{"-R", "--READ"}
var modeL [2]string = [2]string{"-L", "--LENGTH"}

func main () {

    if len(os.Args) < 2 { get_error("args < 2") }

    if 
    strings.ToUpper(os.Args[1]) == modeG[0] ||
    strings.ToUpper(os.Args[1]) == modeG[1] {

        check_args("G", os.Args)

        max, err := strconv.Atoi(os.Args[3])
        check_error(err)

        create_key(os.Args[2], int32(max))

    } else if 
    strings.ToUpper(os.Args[1]) == modeR[0] ||
    strings.ToUpper(os.Args[1]) == modeR[1] {

        check_args("R", os.Args)

        pos, err := strconv.Atoi(os.Args[3])
        check_error(err)

        quan, err := strconv.Atoi(os.Args[4])
        check_error(err)

        read_key(os.Args[2], int32(pos), int32(quan), get_length(os.Args[2]))

    } else if 
    strings.ToUpper(os.Args[1]) == modeL[0] ||
    strings.ToUpper(os.Args[1]) == modeL[1] {

        check_args("L", os.Args)
        fmt.Println("Length of file:", get_length(os.Args[2]))

    } else {
        get_error("mode is not found")
    }
    
    os.Exit(0)
}

func check_args (mode string, args []string) {
    if mode == modeG[0] || mode == modeG[1] {
        if len(args) != 4 {
            get_error("mode = c, args != 4")
        }
    } else if mode == modeR[0] || mode == modeR[1] {
        if len(args) != 5 {
            get_error("mode = r, args != 5")
        }
    } else if mode == modeL[0] || mode == modeL[1] {
        if len(args) != 3 {
            get_error("mode = l, args != 3")
        }
    }
}

func create_key (file_name string, max_num int32) {
    if (max_num < 0) {
        get_error("Max number < 0")
    }
    save_key(file_name, generate_key(max_num))
}

func read_key (file_name string, position, quantity int32, length int32) {
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
