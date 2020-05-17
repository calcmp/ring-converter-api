package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getConversion(w http.ResponseWriter, r *http.Request) {
	log.Println("Get Conversion: Hit")
	w.Header().Set("Content-Type", "application/json")

	var pos InputPos

	requestBody := json.NewDecoder(r.Body).Decode(&pos)
	if requestBody != nil {
		log.Fatal("Error")
		http.Error(w, requestBody.Error(), http.StatusBadRequest)
		return
	}

	log.Println(convertSizes(pos.Position))
	json.NewEncoder(w).Encode(convertSizes(pos.Position))
}

func main() {
	// Init router
	router := mux.NewRouter()

	// Handle requests
	router.HandleFunc("/api/convert", getConversion)
	log.Fatal(http.ListenAndServe(":8081", router))
}
