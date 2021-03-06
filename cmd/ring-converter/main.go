package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getConversion(w http.ResponseWriter, r *http.Request) {
	input := getInput(w, r)
	json.NewEncoder(w).Encode(convertSizes(input.Input))
}

func getMap(w http.ResponseWriter, r *http.Request) {
	input := getInput(w, r)
	json.NewEncoder(w).Encode(returnMap(input.Input))
}

func getInput(w http.ResponseWriter, r *http.Request) Input {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var input Input

	requestBody := json.NewDecoder(r.Body).Decode(&input)
	if requestBody != nil {
		log.Fatal("Error")
		http.Error(w, requestBody.Error(), http.StatusBadRequest)
	}
	return input
}

func main() {
	// Init router
	router := mux.NewRouter()

	// Handle requests
	router.HandleFunc("/api/convert", getConversion).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/map", getMap).Methods("POST", "OPTIONS")
	log.Fatal(http.ListenAndServe(":8081", router))
}
