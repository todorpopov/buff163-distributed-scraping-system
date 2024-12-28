package logger

import (
	"fmt"
	"time"

	"github.com/todorpopov/bdss-common/queue"
)

type Logger struct {
	service string
	queue   queue.QueueGateway
}

func NewLogger(service string, queue queue.QueueGateway) *Logger {
	return &Logger{service, queue}
}

func (log *Logger) Info(msg string) error {
	time := time.Now()

	message := fmt.Sprintf("[%s][%s][INFO] %s", time, log.service, msg)
	return log.queue.Publish("logging", message)
}
