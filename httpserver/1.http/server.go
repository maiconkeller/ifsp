package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "hello world!")
	})

	fmt.Println("starting server... :8080 ")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
