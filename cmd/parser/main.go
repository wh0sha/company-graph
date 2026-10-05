package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"company-graph/internal/model"

	"github.com/PuerkitoBio/goquery"
)

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
	if res.StatusCode != http.StatusOK {
		log.Fatal(res.Status)
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var items []model.Company
	doc.Find("a.news-line-link").Each(func(i int, s *goquery.Selection) {
		var item model.Company

		item.Name = s.Find(".news-line-title").Text()

		item.INN = strconv.Itoa(rand.Intn(9000000000) + 1000000000)
		item.OGRN = strconv.Itoa(rand.Intn(9000000000000) + 1000000000000)
		item.Source = "rbc.ru"

		items = append(items, item)
	})

	if len(items) == 0 {
		log.Fatal("нет данных")
	}
	fmt.Print("записей спаршено: ", len(items), "\n")

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
