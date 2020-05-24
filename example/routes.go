package main

import (
	"net/http"

	"github.com/azeezolaniran2016/gomux"
)

func routes() *gomux.Mux {
	handler := gomux.New()

	handler.HandleFunc("/healthcheck", http.MethodGet, DurationMW(getHealthcheck))

	handler.HandleFunc("/docs", http.MethodGet, DurationMW(getDocs))

	handler.HandleFunc("/docs", http.MethodPost, DurationMW(postDocs))

	handler.HandleFunc("/doc", http.MethodPost, DurationMW(postDoc))

	handler.HandleFunc("/docs/:id", http.MethodGet, DurationMW(getDoc))

	return handler
}

func getHealthcheck(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)

	// we are intentionally ignoring all errors returned since this is for demo purposes
	rw.Write([]byte(`{"status": "ok"}`))
}

func getDocs(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(
		`{"docs": [
			{"id": "one", "title": "Document One"},
			{"id": "two", "title": "Document Two"},
			{"id": "three", "title": "Document Three"}
		]}`))
}

func postDocs(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(`{"msg": "documents created"}`))
}

func postDoc(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(`{"msg": "document created"}`))
}

func getDoc(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(`{"id": "one", "title": "Document One"}}`))
}
