package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"
	"tubes2/crawl"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var maincounter *int

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/scrape", handleScrape).Methods("GET")
	r.HandleFunc("/api/wikipedia", handleWikipediaRequest).Methods("GET")
	r.HandleFunc("/api/idsfunc", handleIDSRequest).Methods("GET")
	r.HandleFunc("/api/bfsfunc", handleBFSRequest).Methods("GET")

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

/* Fungsi IDS */
func handleIDSRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	// Validasi input
	query := r.URL.Query().Get("query")
	query2 := r.URL.Query().Get("query2")
	if query == "" || query2 == "" {
		http.Error(w, "Both query parameters are required", http.StatusBadRequest)
		return
	}

	flag := make(chan bool)
	maincounter = &IDSvisits
	go clock(flag)
	flag <- false

	// Panggil IDSWrapper dengan input yang valid
	resultArticle, visitArticle := IDSWrapper(query, query2)

	flag <- true

	// Buat respons JSON dengan format yang diharapkan
	response := struct {
		Keywords []string `json:"keywords"`
		Number   int      `json:"number"`
	}{
		Keywords: resultArticle[0].trail,
		Number:   visitArticle,
	}

	// Encode respons JSON
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error encoding response JSON", http.StatusInternalServerError)
		return
	}

	// Set header dan kirim respons
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

/* Fungsi BFS */
func handleBFSRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}

	query := r.URL.Query().Get("query")
	query2 := r.URL.Query().Get("query2")
	if query == "" || query2 == "" {
		// http.Error(w, "Both query parameters are required", http.StatusBadRequest)
		// return
		fmt.Println("ERROR QUERY PARAMETER")
	}

	// Panggil fungsi BFS dengan input yang valid
	flag := make(chan bool)
	maincounter = &BFSVisits
	go clock(flag)
	flag <- false
	resultArticle, visitArticle := BFS(query, query2)
	flag <- true

	// Buat respons JSON dengan format yang diharapkan
	response := struct {
		Result []string `json:"keywords"`
		Visit  int      `json:"number"`
	}{
		Result: resultArticle[0].trail,
		Visit:  visitArticle,
	}

	// Encode respons JSON
	responseJSON, err := json.Marshal(response)
	if err != nil {
		// http.Error(w, "Error encoding response JSON", http.StatusInternalServerError)
		fmt.Println("ERROR RESPONSE JSON")
		return
	}

	// Set header dan kirim respons
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
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

	result := crawl.Scrape(query)
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

func clock(flag chan bool) {
	var ms int = 0
	var seconds int = 0
	stop := <-flag
	for !stop {
		if !stop {
			time.Sleep(10 * time.Millisecond)
			ms = ms + 1
		}
		select {
		case newstop := <-flag:
			stop = newstop
		default:
			stop = stop
			if ms/100 > seconds {
				seconds = ms / 100
				fmt.Println(seconds, *maincounter)
			}
		}
	}
}
