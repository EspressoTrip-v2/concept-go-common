package events

import (
	"encoding/json"
	"fmt"
	"github.com/EspressoTrip-v2/concept-go-common/exchange/exchangeNames"
	"github.com/EspressoTrip-v2/concept-go-common/exchange/exchangeTypes"
	"github.com/EspressoTrip-v2/concept-go-common/exchange/queue/queueInfo"
	libErrors "github.com/EspressoTrip-v2/concept-go-common/liberrors"
	"github.com/EspressoTrip-v2/concept-go-common/microservice/microserviceNames"
	"github.com/streadway/amqp"
)

type Publisher interface {
	Publish(data interface{}) *libErrors.CustomError
	failOnError(err error) *libErrors.CustomError
}

type EventPublish struct {
	rabbitConnection *amqp.Connection
	exchangeName     exchangeNames.ExchangeNames
	exchangeType     exchangeTypes.ExchangeType
	queueName        queueInfo.QueueInfo
	publisherName    string
	serviceName      microserviceNames.MicroserviceNames
}

// NewEventPublish creates a new publisher to use
func NewEventPublish(rabbitConnection *amqp.Connection, exchangeName exchangeNames.ExchangeNames, exchangeType exchangeTypes.ExchangeType,
	queueName queueInfo.QueueInfo, publisherName string, serviceName microserviceNames.MicroserviceNames) *EventPublish {

	return &EventPublish{rabbitConnection: rabbitConnection, exchangeName: exchangeName,
		exchangeType: exchangeType, queueName: queueName, publisherName: publisherName, serviceName: serviceName}

}

// Publish publishes the message to RabbitMQ this is usually chained to NewEventPublish
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
	if err := p.failOnError(err); err != nil {
		return err
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
		fmt.Printf("[publisher:%v]: Publisher error: %v | queue:%v | error: %v\n", p.publisherName, p.exchangeName, p.queueName, err.Error())
		return libErrors.NewEventPublisherError(err.Error())
	}
	return nil
}
