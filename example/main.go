package main

import (
	"log"
	"net/http"
)

func main() {
	srv := http.Server{
		Addr:    ":5050",
		Handler: routes(),
	}

	log.Print("Starting server...")

	err := srv.ListenAndServe()

	log.Printf("Error occured from server: %+v", err)
}
