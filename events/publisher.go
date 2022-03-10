package events

import (
	"encoding/json"
	"fmt"
	libErrors "github.com/EspressoTrip-v2/concept-go-common/lib-errors"
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

func (p *EventPublish) Publish(data interface{}) *libErrors.CustomError {
	ch, err := p.rabbitConnection.Channel()
	if err := p.failOnError(err); err != nil {
		return err
	}
	defer ch.Close()

	// declare the exchange if it does not exist
	err = ch.ExchangeDeclare(string(p.exchangeName), string(p.exchangeType), true, false, false, false, nil)
	if err := p.failOnError(err); err != nil {
		return err
	}

	// serialize data
	marshal, err := json.Marshal(data)
	if err != nil {
		return libErrors.NewBadRequestError("Problem serializing data for RabbitMQ message")
	}

	err = ch.Publish(string(p.exchangeName), string(p.queueName), true, false, amqp.Publishing{
		ContentType:  "text/plain",
		DeliveryMode: amqp.Persistent,
		Body:         marshal,
	})
	if err := p.failOnError(err); err != nil {
		return err
	}
	return nil
}

func (p *EventPublish) failOnError(err error) *libErrors.CustomError {
	if err != nil {
		fmt.Printf("[publisher:%v]: Publisher error: %v | queue:%v", p.publisherName, p.exchangeName, p.queueName)
		return libErrors.NewEventPublisherError(err.Error())
	}
	return nil
}
