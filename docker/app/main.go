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
		// res, err := os.ReadFile("/app/mounted/test.txt")
		// if err != nil {

		// }
		fmt.Fprint(w, "hello, world!")
	})
	log.Println(http.ListenAndServe(os.Args[1], nil))
}
