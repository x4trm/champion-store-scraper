package scraper

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ScrapeItems(item string, url string, writer *csv.Writer) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("Error while loading page: %w", err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		return fmt.Errorf("Error parsing HTML: %w", err)
	}

	doc.Find(".details").Each(func(i int, s *goquery.Selection) {
		collection := strings.TrimSpace(s.Find(".product-attribute__collection").Text())
		name := strings.TrimSpace(s.Find(".product-item-name a").Text())
		price := strings.TrimSpace(s.Find(".normal-price .price-wrapper").Text())

		data := []string{item, collection, name, price}
		if err := writeCSV(data, writer); err != nil {
			fmt.Println("Error while writing to CSV:", err)
		}
	})

	return nil
}

func writeCSV(data []string, writer *csv.Writer) error {
	trimmedData := make([]string, len(data))
	for i, cell := range data {
		trimmedData[i] = strings.TrimSpace(cell)
	}
	return writer.Write(trimmedData)
}
