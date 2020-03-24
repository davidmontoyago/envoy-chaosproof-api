package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/davidmontoyago/go-event-ingestor-api/pkg/log"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/hello", hello).Methods("GET")

	srv := &http.Server{
		Handler:      router,
		Addr:         ":9000",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	log.Error.Fatal(srv.ListenAndServe())
}

func hello(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"ok": "true", "hello": "World"})
}
