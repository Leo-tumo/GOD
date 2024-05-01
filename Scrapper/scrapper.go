package Scrapper

type Scrapper interface {
	Scrape(url string) []string
}

func Scrape(urls []string, s Scrapper) {
	for _, url := range urls {
		s.Scrape(url)
	}
}
