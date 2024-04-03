package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	result := scrape("pertamina") // hanya contoh
	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}
}

func scrape(articleName string) []string {

	var links []string

	// creating a new Colly instance
	c := colly.NewCollector()

	// visiting the target page
	c.OnHTML("div#mw-content-text a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if link[:6] == "/wiki/" {
			links = append(links, link[6:])
		}

	})

	c.Visit("https://en.wikipedia.org/wiki/" + articleName)

	return links
}
