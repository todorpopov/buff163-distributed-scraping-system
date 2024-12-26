package scraper

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/todorpopov/bdss-item-scraper/schema"
	"github.com/todorpopov/bdss-item-scraper/utils"
)

type Scraper struct {
	shouldScrape bool
}

func NewScraper() *Scraper {
	return &Scraper{}
}

func (s *Scraper) getJson(itemCode string) *schema.ResponseJson {
	url := utils.GetItemsApi(itemCode)

	response, err := http.Get(url)
	if err != nil {
		return nil
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil
	}

	var data schema.ResponseJson

	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil
	}

	return &data
}

func (s *Scraper) Scrape() {}

func (s *Scraper) StartScraper() bool {
	return true
}
