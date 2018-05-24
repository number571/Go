package main
import "fmt"
 
func main() {
	for i := 1; i < 30; i++ {
		fmt.Println(<-useCH(i))
	}
}

func useCH(x int) chan uint {
	var ch chan uint = make(chan uint)
	go func() {
		ch <- fibonacci(x)
	} ()
	return ch
}

func fibonacci(x int) uint {
	if x <= 3 {
		return 1
	} else {
		return fibonacci(x-1) + fibonacci(x-2)
	}
}
