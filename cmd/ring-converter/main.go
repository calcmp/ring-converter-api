package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getConversion(w http.ResponseWriter, r *http.Request) {
	log.Println("Get Conversion: Hit")

	input := getInput(w, r)

	log.Println(convertSizes(input.Input))

	json.NewEncoder(w).Encode(convertSizes(input.Input))
}

func getMap(w http.ResponseWriter, r *http.Request) {
	log.Println("Get Conversion: Hit")

	input := getInput(w, r)

	log.Println(returnMap(input.Input))

	/*data, err := json.Marshal(returnMap(input.Input))
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(data)

	jsonStr := string(data)
	fmt.Println(jsonStr)*/

	json.NewEncoder(w).Encode(returnMap(input.Input))
}

func getInput(w http.ResponseWriter, r *http.Request) Input {
	log.Println("Get Input: Hit")
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
