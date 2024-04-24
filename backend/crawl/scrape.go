package crawl

import (
	"strings"

	"github.com/gocolly/colly"
)

func Scrape(articleName string) []string {

	var links []string

	// creating a new Colly instance
	c := colly.NewCollector()

	// visiting the target page
	c.OnHTML("div#mw-content-text a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if strings.HasPrefix(link, "/wiki/") && !strings.Contains(link, ":") {
			links = append(links, link[6:])
		}

	})

	c.Visit("https://en.wikipedia.org/wiki/" + articleName)

	return links
}

//!strings.Contains(link, "Special:") && !strings.Contains(link, "File:") && !strings.Contains(link, "Template:") && !strings.Contains(link, "Portal:") && !strings.Contains(link, "Talk:")
