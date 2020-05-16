package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Init router
	router := mux.NewRouter()

	// Handle requests
	router.HandleFunc("/api/convert", getConversion)
	log.Fatal(http.ListenAndServe(":8081", router))
}

func getConversion(w http.ResponseWriter, r *http.Request) {
	log.Println(convertToMillis(20))
}
