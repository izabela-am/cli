package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

// Scrapes HTML page
func HTMLParser(url string) {
	html := downloadHTML(url)
	links := extractLinks(html)
	format(links)

}

func format(links []string) {
	for index, link := range links {
		fmt.Println(index)
		fmt.Println(link)
	}
}

func extractLinks(html []byte) []string {
	var links []string
	html_string := string(html)

	regex := regexp.MustCompile(`(?i)<a[^>]+href=("|')([^"']+)["']`)

	found := regex.FindAllStringSubmatch(html_string, -1)

	for _, match := range found {
		links = append(links, match[2])
	}

	return links
}

func downloadHTML(url string) []byte {
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	html, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	return html
}
