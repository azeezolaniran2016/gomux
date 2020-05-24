package main

import (
	"log"
	"net/http"
	"time"

	"github.com/azeezolaniran2016/gomux"
)

// DurationMW is an example middleware which logs the duration of the request
func DurationMW(hh gomux.HandlerFunc) gomux.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		start := time.Now()
		hh(rw, req)
		log.Printf("request_duration:: method: %s, path: %s, duration: %v", req.Method, req.URL.Path, time.Now().Sub(start))
	}
}
