package main

import (
	"github.com/todorpopov/bdss-common/logger"
	"github.com/todorpopov/bdss-common/queue"
	"github.com/todorpopov/bdss-common/utils"
	"github.com/todorpopov/bdss-item-scraper/src/scraper"
	"github.com/todorpopov/bdss-item-scraper/src/service"
)

func main() {
	queueGateway, err := queue.NewQueueGateway("amqp://guest:guest@localhost:5672/")
	if err != nil {
		utils.FailOnError(err, "An error occured when connectiong to RabbitMQ")
	}

	err = queueGateway.QueueDeclare("logging") // Logging queue declare
	if err != nil {
		utils.FailOnError(err, "Could not declare logging queue!")
	}

	err = queueGateway.QueueDeclare("items") // Items queue declare
	if err != nil {
		utils.FailOnError(err, "Could not declare items queue!")
	}

	logger := logger.NewLogger("ITEM-SCRAPER", *queueGateway)

	scraper := scraper.NewScraper(*queueGateway, *logger)

	service := service.NewService(*scraper)

	service.StartScraping()

	queueGateway.Close()
}
