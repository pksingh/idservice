package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/hello", hello)

	fmt.Println("Server Starting on 8080")
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hi!")
}
