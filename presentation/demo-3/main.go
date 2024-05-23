package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

func main() {
	http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("Error starting the server")
	}
}
