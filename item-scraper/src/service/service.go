package service

import (
	"github.com/todorpopov/bdss-common/logger"
	"github.com/todorpopov/bdss-common/queue"
	"github.com/todorpopov/bdss-item-scraper/src/scraper"
)

type ItemScraperService struct {
	scraper *scraper.Scraper
	logger  *logger.Logger
	qg      *queue.QueueGateway
}

func NewItemScraperService(logger *logger.Logger, qg *queue.QueueGateway) *ItemScraperService {
	scraper := scraper.NewScraper(qg, logger)

	return &ItemScraperService{scraper, logger, qg}
}

func (svc *ItemScraperService) StartScraping() {
	svc.logger.Info("Item Scraping has started!")
	svc.scraper.StartScraping()
}

func (svc *ItemScraperService) StopScraping() {
	svc.logger.Info("Item Scraping has stopped!")
	svc.scraper.StopScraping()
}
