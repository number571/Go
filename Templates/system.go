package main

import (
	"strings"
    "os/exec"
    "fmt"
)

func main() {
	result, err := system("ls -l")
	check_error(err)
	fmt.Println(string(result))
}

func system (line_command string) ([]byte, error) {
    var command []string = strings.Split(line_command, " ")
    return exec.Command(command[0], command[1:]...).Output()
}

func check_error (err error) {
	if err != nil {
		fmt.Println(err)
	}
}
