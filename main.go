package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"project/config"
	"project/scraper"
	"strconv"
)

func main() {

	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		fmt.Println("Error loading config file:", err)
		return
	}

	if _, err := os.Stat("output"); os.IsNotExist(err) {
		os.Mkdir("output", os.ModePerm)
	}

	file, err := os.Create("output/data.csv")
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	_ = writer.Write([]string{"Item", "Collection", "Name", "Sex", "Price", "Product Link", "Image Link"})

	for _, entry := range cfg.Entries {
		for page := 1; page <= entry.MaxPages; page++ {
			url := entry.BaseUrl + strconv.Itoa(page)
			err := scraper.ScrapeItems(entry.Item, entry.Sex, url, writer)
			if err != nil {
				fmt.Println("Error scraping", entry.Item, ":", err)
			}
		}
	}

	fmt.Println("Data saved to file:", "data.csv")
}
