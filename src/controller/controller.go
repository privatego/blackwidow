package controller

import (
	"extractor"
)

func CrawlController() {

}

func Crawl(url string, depth int, extractor Extractor) {
	if depth <= 0 {
		return
	}
	body, urls, err := extractor.Extract(url)
	if err != nil {
		log.Println("fetch fail for url ", url)
	}
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}
