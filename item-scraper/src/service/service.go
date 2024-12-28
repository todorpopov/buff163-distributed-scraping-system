package service

import (
	"github.com/todorpopov/bdss-item-scraper/src/scraper"
)

type Service struct {
	scraper scraper.Scraper
}

func NewService(scraper scraper.Scraper) *Service {
	return &Service{scraper}
}

func (svc *Service) StartScraping() {
	svc.scraper.StartScraping()
}

func (svc *Service) StopScraping() {
	svc.scraper.StopScraping()
}
