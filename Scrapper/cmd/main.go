package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		title := e.Text
		fmt.Printf("Link found: %s -> %s\n", e.Request.AbsoluteURL(link), title)
	})

	c.Visit("https://www.google.com")

}
