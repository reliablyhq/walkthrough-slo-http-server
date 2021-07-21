package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

const (
	paramLatency    = "latency"
	paramStatusCode = "statuscode"
)

func main() {
	log.Println("server starting...")

	var portString string
	if port, err := strconv.Atoi(os.Getenv("PORT")); err == nil {
		portString = fmt.Sprintf(":%d", port)
	} else {
		log.Fatal(err)
	}

	router := httprouter.New()
	router.GET("/", index)

	log.Println("server listening at", portString)
	log.Fatal(http.ListenAndServe(portString, router))
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	log.Println(r.Method, r.URL.String())

	var delay time.Duration
	if l, err := time.ParseDuration(r.URL.Query().Get(paramLatency)); err == nil {
		delay = l
	}

	log.Println("delaying by", delay.String())

	statusCode := 200
	if s, err := strconv.Atoi(r.URL.Query().Get(paramStatusCode)); err == nil {
		statusCode = s
	}

	log.Println("returning status code", statusCode)

	time.Sleep(delay)

	w.WriteHeader(statusCode)
}
