package main
import ("fmt"; "strings")
 
func main() {
	var (
		mode string
		key int8
		message string
	)

	fmt.Scan(&mode, &key, &message)
	mode = strings.ToUpper(mode)

	if mode != "E" && mode != "D" {
		panic("Error: mode is not found")
	}

	fmt.Println("Final message:", caesar(mode, key, message))
}

func caesar(mode string, key int8, message string) string {
	if mode == "E" {
		return encrypt(key, message)
	} else {
		return decrypt(key, message)
	}
}

func encrypt(key int8, message string) (encrypt string) {
	for _, symbol := range strings.ToUpper(message) {
		encrypt += string((int8(symbol) + key - 13) % 26 + 'A')
	}; return
}

func decrypt(key int8, message string) (decrypt string) {
	for _, symbol := range strings.ToUpper(message) {
		decrypt += string((int8(symbol) - key - 13) % 26 + 'A')
	}; return
}
