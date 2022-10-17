package amqp

import (
	"log"

	"github.com/rabbitmq/amqp091-go"
)

func CreateConnection(amqpUrl string) *amqp091.Connection {
	conn, err := amqp091.Dial(amqpUrl)

	if err != nil {
		log.Fatalf("Could not open AMQP connection: %s", err)
	}

	AssertTopics(conn)

	return conn
}
