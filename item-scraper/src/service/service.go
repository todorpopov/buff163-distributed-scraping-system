package service

import (
	"github.com/todorpopov/bdss-common/logger"
	"github.com/todorpopov/bdss-item-scraper/src/scraper"
)

type Service struct {
	scraper scraper.Scraper
	logger  logger.Logger
}

func NewService(scraper scraper.Scraper, logger logger.Logger) *Service {
	return &Service{scraper, logger}
}
