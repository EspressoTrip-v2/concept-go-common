package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
)

type Rabbiter interface {
	Connect() *amqp.Connection
}

type RabbitClient struct {
	rabbitUrl  string
	clientName string
}

func (r RabbitClient) Connect() *amqp.Connection {
	conn, err := amqp.Dial(r.rabbitUrl)
	if err != nil {
		fmt.Printf("[rabbit:%v]: Failed to connect", r.clientName)
	}
	defer conn.Close()
	return conn
}

func NewRabbitClient(rabbitUrl string, clientName string) *RabbitClient {
	return &RabbitClient{rabbitUrl: rabbitUrl, clientName: clientName}
}
