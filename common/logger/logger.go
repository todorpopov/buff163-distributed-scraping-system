package logger

import (
	"fmt"
	"time"

	"github.com/todorpopov/bdss-common/queue"
	"github.com/todorpopov/bdss-common/shared"
)

type Logger struct {
	service string
	qg      queue.QueueGateway
}

func NewLogger(service string, qg queue.QueueGateway) *Logger {
	err := qg.QueueDeclare("logging") // Logging queue declare
	if err != nil {
		shared.FailOnError(err, "Could not declare logging queue!")
	}

	return &Logger{service, qg}
}

func (log *Logger) Info(msg string) error {
	time := time.Now().Format("2006-01-02 15:04:05")

	message := fmt.Sprintf("[%s][%s][INFO] %s", time, log.service, msg)
	return log.qg.Publish("logging", message)
}

func (log *Logger) Debug(msg string) error {
	time := time.Now().Format("2006-01-02 15:04:05")

	message := fmt.Sprintf("[%s][%s][DEBUG] %s", time, log.service, msg)
	return log.qg.Publish("logging", message)
}

func (log *Logger) Error(err error, msg string) error {
	time := time.Now().Format("2006-01-02 15:04:05")

	message := fmt.Sprintf("[%s][%s][ERROR] %s: %s", time, log.service, msg, err.Error())
	return log.qg.Publish("logging", message)
}
