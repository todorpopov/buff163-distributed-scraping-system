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
	time := time.Now().Format("2006-01-02 15:04:05")

	message := fmt.Sprintf("[%s][%s][INFO] %s", time, log.service, msg)
	return log.queue.Publish("logging", message)
}

func (log *Logger) Debug(msg string) error {
	time := time.Now().Format("2006-01-02 15:04:05")

	message := fmt.Sprintf("[%s][%s][DEBUG] %s", time, log.service, msg)
	return log.queue.Publish("logging", message)
}

func (log *Logger) Error(err error, msg string) error {
	time := time.Now().Format("2006-01-02 15:04:05")

	message := fmt.Sprintf("[%s][%s][ERROR] %s: %s", time, log.service, msg, err.Error())
	return log.queue.Publish("logging", message)
}
