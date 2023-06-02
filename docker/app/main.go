package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	log.Printf("Service is listening... [%s]\n", os.Args[1])
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		res, err := os.ReadFile("/mounted/test.txt")
		if err != nil {
			fmt.Fprintf(w, "error: %s", err.Error())
		}
		fmt.Fprint(w, string(res))
	})
	log.Println(http.ListenAndServe(os.Args[1], nil))
}
