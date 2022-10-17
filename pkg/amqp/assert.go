package amqp

import (
	"encoding/json"
	"log"
	"os"
	"path"

	"github.com/rabbitmq/amqp091-go"
	"github.com/rellyson/gobet/config"
)

func AssertTopics(conn *amqp091.Connection) {
	rootPath, _ := os.Getwd()
	jsonFile, _ := os.ReadFile(path.Join(
		rootPath,
		"config",
		"amqp-topics.json",
	))

	var topics []config.AmqpTopic
	json.Unmarshal(jsonFile, &topics)

	ch, err := conn.Channel()

	if err != nil {
		log.Fatalf("Cannot assert AMQP topics: %s", err)
	}

	for _, topic := range topics {
		ch.ExchangeDeclare(
			topic.Name,
			"topic",
			true,
			false,
			false,
			false,
			nil,
		)

		ch.QueueDeclare(
			topic.QueueName,
			true,
			false,
			false,
			false,
			nil,
		)

		ch.QueueBind(
			topic.QueueName,
			topic.RoutingKey,
			topic.Name,
			false,
			nil,
		)

		log.Printf("Assert topic %s -> %s (%s)",
			topic.Name,
			topic.QueueName,
			topic.RoutingKey,
		)
	}

	defer ch.Close()
}
