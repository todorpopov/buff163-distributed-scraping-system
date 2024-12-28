package scraper

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/todorpopov/bdss-common/logger"

	"github.com/todorpopov/bdss-item-scraper/src/schema"
	"github.com/todorpopov/bdss-item-scraper/src/utils"
)

type Scraper struct {
	shouldScrape bool
	timeout      int
	logger       logger.Logger
	itemCodes    []string
}

func NewScraper(logger logger.Logger) *Scraper {
	itemCodes := utils.ParseItemCodesFile()
	return &Scraper{true, 1, logger, itemCodes}
}

func (s *Scraper) getSerializedDtos(itemCode string) *[]schema.ItemDTO {
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

	return data.ToItemDtos()
}

func (s *Scraper) serializeDtos(dtos []schema.ItemDTO) *[]string {
	var serializedDtos []string
	for i := 0; i < len(dtos); i++ {
		dto := dtos[i].Serialize()

		if dto != "" {
			serializedDtos = append(serializedDtos, dto)
		}
	}

	return &serializedDtos
}
