package main

import (
	"github.com/todorpopov/bdss-common/queue"
	"github.com/todorpopov/bdss-common/utils"
)

func main() {
	queueGateway, err := queue.NewQueueGateway("amqp://guest:guest@localhost:5672/")
	if err != nil {
		utils.FailOnError(err, "An error occured when connectiong to RabbitMQ")
	}

	err = queueGateway.QueueDeclare("logging") // Logging Queue declare
	if err != nil {
		utils.FailOnError(err, "Could not declare queue!")
	}

	// logger := logger.NewLogger("ITEM-SCRAPER", *queueGateway)

	// scraper := scraper.NewScraper(*logger)
}
