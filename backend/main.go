package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// func main() {
// 	http.HandleFunc("/scrape", handleScrape)
// 	http.HandleFunc("/api/wikipedia", handleWikipediaRequest)

// 	fmt.Println("server listening on port 8080...")
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/scrape", handleScrape).Methods("GET")
	r.HandleFunc("/api/wikipedia", handleWikipediaRequest).Methods("GET")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	fmt.Println("server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

/* Fungsi Scrape */
func handleScrape(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

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

/* Fungsi Menampilkan Hasil Pencarian Dari Wikipedia API */
func handleWikipediaRequest(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
			http.Error(w, "Query parameter is required", http.StatusBadRequest)
			return
	}

	wikipediaURL := "https://en.wikipedia.org/w/api.php?action=query&format=json&list=search&srsearch=" + url.QueryEscape(query)
	response, err := http.Get(wikipediaURL)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}
	defer response.Body.Close()

	var data interface{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}