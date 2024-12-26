package main

import (
	// "github.com/todorpopov/bdss-item-scraper/scraper"
	// "github.com/todorpopov/bdss-item-scraper/service"
	"github.com/todorpopov/bdss-item-scraper/system"
)

func main() {
	// scraper := scraper.NewScraper()

	// service := service.NewService(*scraper)

	// service.StartScraper()

	fileHandler := system.NewFileHandler("item-scraper/system/item-codes.txt")
	fileHandler.ParseFile()
}
