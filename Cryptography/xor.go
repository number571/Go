package main
import "fmt"
 
func main() {
	var (
		key uint
		message string
	)
	fmt.Scan(&key, &message)
	fmt.Println("Final message:", encryptDecrypt(key, message))
}

func encryptDecrypt(key uint, message string) (final string) {
	for _, symbol := range message {
		final += string(uint(symbol) ^ key)
	}; return
}
