package main

import (
    "fmt"
    "os"
    "net/http"
)

const BUFF = 512

func main() {
    fmt.Println(urlopen("https://google.com"))
}

func urlopen(url string) string {
    var (
        resp *http.Response
        html_page string
        buffer []byte
        length int
        err error
    )
    resp, err = http.Get(url) 
    check_error(err)
    defer resp.Body.Close()

    buffer = make([]byte, BUFF)
    for {
        length, err = resp.Body.Read(buffer)
        html_page += string(buffer[:length])
        if length == 0 || err != nil{ break }
    }
    return html_page
}

func check_error (err error) {
    if err != nil {
        fmt.Println("Error:", err)
        os.Exit(1)
    }
}
