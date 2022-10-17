package config

type AmqpTopic struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	RoutingKey string `json:"routingKey"`
	QueueName  string `json:"queueName"`
}
