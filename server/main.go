package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

const (
	paramLatency    = "latency"
	paramStatusCode = "statuscode"
)

func main() {
	router := httprouter.New()

	router.GET("/", index)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var delay time.Duration
	if l, err := time.ParseDuration(r.URL.Query().Get(paramLatency)); err != nil {
		delay = l
	}

	statusCode := 200
	if s, err := strconv.Atoi(r.URL.Query().Get(paramStatusCode)); err != nil {
		statusCode = s
	}

	time.Sleep(delay)

	w.WriteHeader(statusCode)
}
