package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/scrape", handleScrape)
	fmt.Println("server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleScrape(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		// Respond to preflight requests
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // Allow requests from this origin
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	// Handle GET request for data scraping
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "Query parameter is required", http.StatusBadRequest)
		return
	}

	result := scrape(query)
	responseJSON, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // Allow requests from this origin
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}
