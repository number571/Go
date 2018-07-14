package main

import (
    "os"
    "fmt"
    "regexp"
    "net/http"
)

const BUFF = 512

func main() {
    var html = urlopen("https://google.com")
    for _, link := range get_links(html) {
        fmt.Println(link)
    }
}

func get_links (html string) []string {
    var (
        result [][]string
        slice_links []string
        regular *regexp.Regexp 
    )

    regular = regexp.MustCompile(`href\s*=\s*['\"]([^'\"]+)`)
    result = regular.FindAllStringSubmatch(html, -1)
    
    for _, link := range result {
        slice_links = append(slice_links, link[1])
    }
    return slice_links
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
