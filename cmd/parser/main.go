package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Item struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

func main() {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("GET", "https://www.rbc.ru", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/115.0.0.0 Safari/537.36 Edg/115.0.0.0")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var items []Item
	doc.Find("a.news-line-link").Each(func(i int, s *goquery.Selection) {
		var item Item

		item.Title = s.Find(".news-line-title").Text()

		if href, exists := s.Attr("href"); exists {
			item.Link = href
		}

		items = append(items, item)
	})

	jsonData, err := json.MarshalIndent(items, "", "    ")
	if err != nil {
		log.Fatal(err)
	}

	filename := "result.json"
	err = os.WriteFile(filename, jsonData, 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("over")

}
