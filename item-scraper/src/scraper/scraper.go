package scraper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/todorpopov/bdss-common/logger"
	"github.com/todorpopov/bdss-common/queue"

	"github.com/todorpopov/bdss-item-scraper/src/schema"
	"github.com/todorpopov/bdss-item-scraper/src/utils"
)

type Scraper struct {
	shouldScrape bool
	timeout      int
	itemCodes    []string
	qg           queue.QueueGateway
	logger       logger.Logger
}

func NewScraper(qg queue.QueueGateway, logger logger.Logger) *Scraper {
	itemCodes := utils.ParseItemCodesFile()
	return &Scraper{true, 1, itemCodes, qg, logger}
}

func (s *Scraper) getItemDtos(itemCode string) *[]schema.ItemDTO {
	url := utils.GetItemsApi(itemCode)

	response, err := http.Get(url)
	if err != nil {
		s.logger.Error(err, "Could not fetch URL!")
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		s.logger.Error(err, "Could not fetch URL!")
	}

	var data schema.ResponseJson

	err = json.Unmarshal(body, &data)
	if err != nil {
		s.logger.Error(err, "Could not fetch URL!")
	}

	dtos := data.ToItemDtos()

	successMsg := fmt.Sprintf("Succesfully managed to scrape item code [%s] - %s items total", itemCode, strconv.Itoa(len(*dtos)))
	s.logger.Info(successMsg)

	return dtos
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

func (s *Scraper) publishToQueue(serializedDtos []string) {
	for i := 0; i < len(serializedDtos); i++ {
		if serializedDtos[i] != "" {
			err := s.qg.Publish("items", serializedDtos[i])

			if err != nil {
				s.logger.Error(err, "Error occured when publishing item to queue!")
			}
		}
	}
}

func (s *Scraper) StartScraping() {
	s.shouldScrape = true

	for i := 0; i < len(s.itemCodes); i++ {
		if !s.shouldScrape {
			break
		}

		dtos := s.getItemDtos(s.itemCodes[i])
		serializedDtos := s.serializeDtos(*dtos)
		s.publishToQueue(*serializedDtos)

		time.Sleep(5 * time.Second)
	}
}

func (s *Scraper) StopScraping() {
	s.shouldScrape = false
}
