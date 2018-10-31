package main

/*
int factorial(int x) {
    int s = 1;
    while (x != 0) {
        s *= x;
        --x;
    }
    return s;
}
*/
import "C"

import "fmt"

func main() {
    fmt.Println(C.factorial(6))
}
