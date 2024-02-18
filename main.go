package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var foundLinks = make(map[string]bool)

func makeRequest(url string) *goquery.Document {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil
	}
	defer response.Body.Close()

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		fmt.Println("Error parsing response body:", err)
		return nil
	}
	return doc
}

func crawl(url string) {
	doc := makeRequest(url)
	if doc == nil {
		return
	}
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		foundLink, _ := s.Attr("href")
		if foundLink != "" {
			if strings.Contains(foundLink, "#") {
				foundLink = strings.Split(foundLink, "#")[0]
			}
			if strings.Contains(foundLink, url) && !foundLinks[foundLink] {
				foundLinks[foundLink] = true
				fmt.Println(foundLink)
				// recursive
				crawl(foundLink)
			}
		}
	})
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter target URL: ")
	targetURL, _ := reader.ReadString('\n')
	targetURL = strings.TrimSpace(targetURL)
	crawl(targetURL)
}
