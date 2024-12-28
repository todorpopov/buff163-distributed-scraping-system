package main

import (
	"github.com/todorpopov/bdss-common/logger"
	"github.com/todorpopov/bdss-common/queue"
	"github.com/todorpopov/bdss-common/shared"
	"github.com/todorpopov/bdss-item-scraper/src/service"
)

func main() {
	queueGateway, err := queue.NewQueueGateway("amqp://guest:guest@localhost:5672/")
	if err != nil {
		shared.FailOnError(err, "An error occured when connectiong to RabbitMQ!")
	}

	logger := logger.NewLogger("ITEM-SCRAPER", *queueGateway)

	svc := service.NewItemScraperService(logger, queueGateway)
	svc.StartScraping()

	queueGateway.Close()
}
