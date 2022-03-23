package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"math/rand"
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
	http.HandleFunc("/health", GetHealth)
	http.HandleFunc("/idgen", GetIdgen)

	// fmt.Println("Server Starting on 8080")
	log.Println(fmt.Sprintf("idservice Running on port %v ...", *port))

	// http.ListenAndServe(":8080", nil)
	log.Fatal(http.ListenAndServe(
		fmt.Sprintf(":%v", *port), nil))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "idservice > hi!")
}

func GetHealth(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"status\": \"%s\"}", status)
		// fmt.Fprintf(w, "{\"status\": \"%s\", \"error\": \"%v\"}", status, err.Error())
	} else {
		fmt.Fprintf(w, "{\"status\": \"ok\"}")
	}
}

func GetIdgen(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "{\"error\": \"%s\"}", status)
	} else {
		uid := rand.Int()

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "{\"uid\": \"%v\"}", uid)
	}

}
