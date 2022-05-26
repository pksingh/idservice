package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	snowid "github.com/pksingh/idservice/snowid"
)

const (
	StatusOK     = "ok"
	StatusFAILED = "failed"

	StatusInit    = "node initiated"
	StatusNotInit = "node not initiated"
)

var err = errors.New(StatusNotInit)
var status = StatusNotInit

var nStartTime time.Time
var nId int64
var nTimeBits int8
var nNodeBits int8
var nCountBits int8

func main() {
	port := flag.Int("p", 80, "Port")
	flag.Parse()

	InitNode(0, time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC), 42, 5, 16)

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/health", GetHealth)
	http.HandleFunc("/idgen", GetIdgen)
	http.HandleFunc("/idmeta", GetIdmeta)
	http.HandleFunc("/parseid", GetIdparsed)

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
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "{\"error\": \"%s\"}", status)
	} else {
		uid := snowid.NextId()

		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "{\"uid\": \"%v\"}", uid)
	}

}

func GetIdmeta(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{\"start_time\": \"%s\", \"node_id\": %d, \"time_bits\": %d, \"node_bits\": %d, \"count_bits\": %d}", nStartTime, nId, nTimeBits, nNodeBits, nCountBits)
}

func GetIdparsed(w http.ResponseWriter, r *http.Request) {
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "{\"error\": \"%s\"}", status)
	} else {
		uidStr := r.URL.Query()["uid"]
		uid, err := strconv.ParseUint(uidStr[0], 10, 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "{\"error\": \"%s\"}", err.Error())
			return
		}
		sid := snowid.ParseId(uid)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "{\"id\": \"%d\", \"time\": %d, \"nodeId\": %d, \"sequence\": %d}", sid.ID, sid.Timestamp, sid.NodeId, sid.Sequence)
	}
}

func InitNode(id int64, startTime time.Time, timeBits, nodeBits, countBits int64) {
	nId = id
	nTimeBits = int8(timeBits)
	nNodeBits = int8(nodeBits)
	nCountBits = int8(countBits)
	nStartTime = startTime
	err = snowid.SetNode(int64(id), startTime, int64(timeBits), int64(nodeBits), int64(countBits))
	if err != nil {
		status = StatusFAILED
	}
}

func InitDefaultNode() {
	nId = 0
	nTimeBits = 42
	nNodeBits = 5
	nCountBits = 16
	nStartTime = time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)

	err = snowid.SetNode(int64(nId), nStartTime, int64(nTimeBits), int64(nNodeBits), int64(nCountBits))
	if err != nil {
		status = StatusFAILED
	}
}
