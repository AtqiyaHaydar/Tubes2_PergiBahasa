package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"tubes2/crawl"
)

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

// /* Fungsi Menampilkan Hasil Pencarian Dari Wikipedia API */
// func handleWikipediaRequest(w http.ResponseWriter, r *http.Request) {
// 	query := r.URL.Query().Get("query")
// 	if query == "" {
// 			http.Error(w, "Query parameter is required", http.StatusBadRequest)
// 			return
// 	}

// 	wikipediaURL := "https://en.wikipedia.org/w/api.php?action=query&format=json&list=search&srsearch=" + url.QueryEscape(query)
// 	response, err := http.Get(wikipediaURL)
// 	if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 	}
// 	defer response.Body.Close()

// 	var data interface{}
// 	err = json.NewDecoder(response.Body).Decode(&data)
// 	if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 	}

// 	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(data)
// }

/* Fungsi Menampilkan Hasil Pencarian Dari Wikipedia API */
func handleWikipediaRequest(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query().Get("query")
	if query == "" {
			http.Error(w, "Query parameter is required", http.StatusBadRequest)
			return
	}

	// Membuat permintaan untuk mendapatkan hasil pencarian dari Wikipedia API
	wikipediaURL := "https://en.wikipedia.org/w/api.php?action=query&format=json&list=search&srsearch=" + url.QueryEscape(query)
	response, err := http.Get(wikipediaURL)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}
	defer response.Body.Close()

	var searchData interface{}
	err = json.NewDecoder(response.Body).Decode(&searchData)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}

	// Menemukan halaman pertama dari hasil pencarian untuk mendapatkan informasi tambahan, termasuk URL gambar
	searchResult := searchData.(map[string]interface{})["query"].(map[string]interface{})["search"].([]interface{})
	if len(searchResult) == 0 {
			// Tidak ada hasil pencarian yang ditemukan
			http.Error(w, "No search results found", http.StatusNotFound)
			return
	}

	// Mengambil judul artikel Wikipedia untuk permintaan lanjutan
	pageTitle := searchResult[0].(map[string]interface{})["title"].(string)

	// Membuat permintaan untuk mendapatkan informasi tambahan tentang halaman artikel, termasuk URL gambar
	pageURL := fmt.Sprintf("https://en.wikipedia.org/w/api.php?action=query&format=json&prop=pageimages&titles=%s&pithumbsize=300", url.QueryEscape(pageTitle))
	pageResponse, err := http.Get(pageURL)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}
	defer pageResponse.Body.Close()

	var pageData interface{}
	err = json.NewDecoder(pageResponse.Body).Decode(&pageData)
	if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
	}

	// Menggabungkan data hasil pencarian dengan informasi tambahan (termasuk URL gambar)
	pageDataMap := pageData.(map[string]interface{})["query"].(map[string]interface{})["pages"].(map[string]interface{})
	for _, v := range pageDataMap {
			searchData.(map[string]interface{})["query"].(map[string]interface{})["search"].([]interface{})[0].(map[string]interface{})["thumbnail"] = v.(map[string]interface{})["thumbnail"]
			break // Hanya memproses halaman pertama
	}

	// Mengatur header dan mengembalikan hasil JSON yang disempurnakan
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(searchData)
}
