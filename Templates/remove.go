package main 

import (
    "fmt"
)

func main() {
    var names []string = []string{"Alice","Bob","Eve","Tom","John"}
    names = remove(names, 3)
    fmt.Println(names)
}

func remove(list []string, num uint32) []string {
    return append(list[:num],list[num+1:]...)
}
