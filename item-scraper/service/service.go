package service

import "github.com/todorpopov/bdss-item-scraper/scraper"

type Service struct {
	scraper scraper.Scraper
}

func NewService(scraper scraper.Scraper) *Service {
	return &Service{scraper}
}

func (s *Service) StartScraper() {
	s.scraper.StartScraper()
}
