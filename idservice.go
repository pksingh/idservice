package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
)

const (
	StatusOK     = "ok"
	StatusFAILED = "failed"

	StatusInit    = "node initiated"
	StatusNotInit = "node not initiated"
)

var status string
var err error

func main() {
	port := flag.Int("p", 80, "Port")
	flag.Parse()

	err = errors.New(StatusNotInit)
	status = StatusNotInit

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/health", health)
	http.HandleFunc("/", health)

	// fmt.Println("Server Starting on 8080")
	log.Println(fmt.Sprintf("idservice Running on port %v ...", *port))

	// http.ListenAndServe(":8080", nil)
	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%v", *port), nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "idservice > hi!")
}

func health(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"status\": \"%s\"}", status)
	} else {
		fmt.Fprintf(w, "{\"status\": \"ok\"}")
	}
}
