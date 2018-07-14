package main

import (
    "fmt"
    "os"
    "net/http"
)

const BUFF = 512

func main() {

    var (
        resp *http.Response
        html_page string
        buffer []byte
        length int
        err error
    )

    resp, err = http.Get("https://google.com") 
    check_error(err)
    defer resp.Body.Close()

    buffer = make([]byte, BUFF)
    for {
        length, err = resp.Body.Read(buffer)
        html_page += string(buffer[:length])
        if length == 0 || err != nil{ break }
    }
    fmt.Println(html_page)
    os.Exit(0)
}

func check_error (err error) {
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
}
