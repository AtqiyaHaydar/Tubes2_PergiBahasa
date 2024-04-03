package main

import (
	"net/http"
	"github.com/gocolly/colly"
	"github.com/PuerkittoBio/goquery"
	"encoding/json"
)

func main() {
	http.HandleFunc("/api/scape", handleScrape)
	http.ListenAndServe(":8080", nil) // Ini Tergantung Port Kira
}

func handleScrape(w http.ResponseWriter, r *http.Request) {}