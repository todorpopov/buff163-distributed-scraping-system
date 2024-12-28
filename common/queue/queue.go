package queue

import (
	"sync"

	"github.com/rabbitmq/amqp091-go"
)

type QueueGateway struct {
	conn    *amqp091.Connection
	channel *amqp091.Channel
	mutex   *sync.Mutex
}

func NewQueueGateway(url string) (*QueueGateway, error) {
	conn, err := amqp091.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	return &QueueGateway{conn, ch, &sync.Mutex{}}, nil
}

func (qg *QueueGateway) Close() {
	defer qg.channel.Close()
	defer qg.conn.Close()
}

func (qg *QueueGateway) QueueDeclare(queueName string) error {
	qg.mutex.Lock()

	defer qg.mutex.Unlock()

	_, err := qg.channel.QueueDeclare(queueName, false, false, false, false, nil)

	return err
}

func (qg *QueueGateway) Publish(queueName string, message string) error {
	qg.mutex.Lock()

	defer qg.mutex.Unlock()

	err := qg.channel.Publish(
		"",
		queueName,
		false,
		false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)

	return err
}
