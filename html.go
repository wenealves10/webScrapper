package webScrapper

import (
	"io/ioutil"
	"net/http"
	"regexp"
)

// Titles is titles
func Titles(urls ...string) <-chan string {
	channel := make(chan string)

	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := ioutil.ReadAll(resp.Body)

			r, _ := regexp.Compile(`<title>(.*?)<\/title>`)
			channel <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}

	return channel
}
