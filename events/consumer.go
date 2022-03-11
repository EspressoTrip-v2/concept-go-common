package events

import (
	"fmt"
	"github.com/EspressoTrip-v2/concept-go-common/exchange/exchangeNames"
	"github.com/EspressoTrip-v2/concept-go-common/exchange/exchangeTypes"
	"github.com/EspressoTrip-v2/concept-go-common/exchange/queue/queueInfo"
	"github.com/EspressoTrip-v2/concept-go-common/microservice/microserviceNames"
	"github.com/streadway/amqp"
)

type EventConsumer struct {
	rabbitConnection *amqp.Connection
	exchangeName     exchangeNames.ExchangeNames
	exchangeType     exchangeTypes.ExchangeType
	queueName        queueInfo.QueueInfo
	consumerName     string
	serviceName      microserviceNames.MicroserviceNames
}

func NewEventConsumer(rabbitConnection *amqp.Connection, exchangeName exchangeNames.ExchangeNames, exchangeType exchangeTypes.ExchangeType,
	queueName queueInfo.QueueInfo, consumerName string, serviceName microserviceNames.MicroserviceNames) *EventConsumer {
	return &EventConsumer{rabbitConnection: rabbitConnection, exchangeName: exchangeName, exchangeType: exchangeType,
		queueName: queueName, consumerName: consumerName, serviceName: serviceName}
}

// Listen subscribes consumer to the set queue {key string} is an optional setting for routing if required
func (c *EventConsumer) Listen(key string) {
	var k string
	if key != "" {
		k = key
	}
	// connect to the rabbit instance
	ch, err := c.rabbitConnection.Channel()
	c.failOnError(err)

	// declare exchange if not exists
	err = ch.ExchangeDeclare(string(c.exchangeName), string(c.exchangeType), true, false, false, false, nil)
	c.failOnError(err)
	defer ch.Close() // TODO: CHECK THIS

	// declare a queue if not exists
	q, err := ch.QueueDeclare(string(c.queueName), true, false, true, false, nil)
	c.failOnError(err)
	err = ch.QueueBind(q.Name, k, string(c.exchangeName), false, nil)
	c.failOnError(err)

	msg, err := ch.Consume(q.Name, "", false, false, false, false, nil)

	// this runs a go routine to acknowledge any incoming message
	go func() {
		for delivery := range msg {
			err := delivery.Ack(false)
			if err != nil {
				fmt.Printf("[consumer:%v] Failed to acknowledge message on: %v | queue:%v | msg: %v", c.exchangeName, c.queueName, delivery.Body)
			}
		}
	}()

}

func (c *EventConsumer) failOnError(err error) {
	if err != nil {
		fmt.Printf("[consumer:%v]: Publisher error: %v | queue:%v", c.consumerName, c.exchangeName, c.queueName)
	}
}
