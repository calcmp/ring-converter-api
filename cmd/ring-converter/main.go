package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage hit")
}

func main() {
	// Init router
	r := mux.NewRouter()

	// Handle requests
	r.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8081", r))
}
