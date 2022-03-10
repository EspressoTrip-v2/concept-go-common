package events

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
)

type EventPublish struct {
	rabbitConnection *amqp.Connection
	exchangeName     ExchangeNames
	exchangeType     ExchangeType
	queueName        QueueInfo
	publisherName    string
	serviceName      MicroserviceNames
}

func NewEventPublish(rabbitConnection *amqp.Connection, exchangeName ExchangeNames, exchangeType ExchangeType,
	queueName QueueInfo, publisherName string, serviceName MicroserviceNames) *EventPublish {

	return &EventPublish{rabbitConnection: rabbitConnection, exchangeName: exchangeName,
		exchangeType: exchangeType, queueName: queueName, publisherName: publisherName, serviceName: serviceName}

}

func (p *EventPublish) Publish(data interface{}) {
	ch, err := p.rabbitConnection.Channel()
	p.failOnError(err)
	defer ch.Close()

	// declare the exchange if it does not exist
	err = ch.ExchangeDeclare(string(p.exchangeName), string(p.exchangeType), true, false, false, false, nil)
	p.failOnError(err)
	// serialize data
	marshal, err := json.Marshal(data)
	p.failOnError(err)

	err = ch.Publish(string(p.exchangeName), string(p.queueName), true, false, amqp.Publishing{
		ContentType:  "text/plain",
		DeliveryMode: amqp.Persistent,
		Body:         marshal,
	})
	p.failOnError(err)
}

func (p *EventPublish) failOnError(err error) {
	if err != nil {
		fmt.Printf("[publisher:%v]: Publisher error: %v | queue:%v", p.publisherName, p.exchangeName, p.queueName)
	}
}
