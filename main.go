package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/health", health).Methods("GET")
	router.HandleFunc("/hello", hello).Methods("GET")

	srv := &http.Server{
		Handler:      router,
		Addr:         ":9000",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func hello(w http.ResponseWriter, r *http.Request) {
	log.Println("headers:")
	for name, values := range r.Header {
		for _, value := range values {
			fmt.Println(name, value)
		}
	}

	json.NewEncoder(w).Encode(map[string]interface{}{"ok": true, "hello": "World"})
}

func health(w http.ResponseWriter, r *http.Request) {

	// check dependencies...

	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
