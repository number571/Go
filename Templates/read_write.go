package main 
import (
	"fmt"
	"os"
)
func main() {
	create("file.txt")
	input("file.txt","Hello World!")
	var message string = ouput("file.txt")
	fmt.Println(message)
}

func create(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
}

func input(filename, message string) {
	file, err := os.OpenFile(
		filename, 
		os.O_WRONLY|os.O_APPEND,
		0666,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	file.WriteString(message)
}

func ouput(filename string) string {
	file, err := os.OpenFile(
		filename,
		os.O_RDONLY,
		0666,
	)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer file.Close()
	var message string
	buffer := make([]byte, 64)
	for {
		m, err := file.Read(buffer)
		if err != nil {
			break
		}
		message += string(buffer[:m])
	}
	return message
}
