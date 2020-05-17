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

	var size Size

	requestBody := json.NewDecoder(r.Body).Decode(&size)
	if requestBody != nil {
		log.Fatal("Error")
		http.Error(w, requestBody.Error(), http.StatusBadRequest)
		return
	}

	log.Println(size.Millis)
	log.Println(size.Cms)

	//size, err := strconv.Atoi(requestBody)

	json.NewEncoder(w).Encode(size.Millis)
	json.NewEncoder(w).Encode(size.Cms)
}

func main() {
	// Init router
	router := mux.NewRouter()

	// Handle requests
	router.HandleFunc("/api/convert", getConversion)
	log.Fatal(http.ListenAndServe(":8081", router))
}
