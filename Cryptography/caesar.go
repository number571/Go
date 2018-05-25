package main
import ("fmt"; "strings")
 
func main() {
	var (
		mode string; key int8
		message string
	)

	fmt.Scan(&mode, &key, &message)
	mode = strings.ToUpper(mode)

	if mode != "E" && mode != "D" {
		panic("Error: mode is not found")
	} else {
		fmt.Println("Final message:", encryptDecrypt(mode, key, message))
	}
}

func encryptDecrypt(mode string, key int8, message string) (final string) {
	if mode == "E" {
		for _, symbol := range strings.ToUpper(message) {
			final += string((int8(symbol) + key - 13) % 26 + 'A')
		}; return
	} else {
		for _, symbol := range strings.ToUpper(message) {
			final += string((int8(symbol) - key - 13) % 26 + 'A')
		}; return
	}
}
