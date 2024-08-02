package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"project/scraper"
	"strconv"
)

func main() {

	file, err := os.Create("data.csv")
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	_ = writer.Write([]string{"Item", "Collection", "Name", "Price"})

	baseUrlHoodie := "https://www.championstore.com/en_gb/men/clothing/hoodies?p="
	maxPagesHoodie := 3

	for page := 1; page <= maxPagesHoodie; page++ {
		url := baseUrlHoodie + strconv.Itoa(page)
		err := scraper.ScrapeItems("Hoodie", url, writer)
		if err != nil {
			fmt.Println("Error scraping hoodies:", err)
		}
	}

	baseUrlTshirt := "https://www.championstore.com/en_gb/men/clothing/t-shirts?p="
	maxPagesTshirt := 5

	for page := 1; page <= maxPagesTshirt; page++ {
		url := baseUrlTshirt + strconv.Itoa(page)
		err := scraper.ScrapeItems("t-shirt", url, writer)
		if err != nil {
			fmt.Println("Error scraping t-shirts:", err)
		}
	}

	fmt.Println("Data saved to file:", "data.csv")
}
