package events

import (
	"fmt"
	"github.com/EspressoTrip-v2/concept-go-common/exchange/exchangeNames"
	"github.com/EspressoTrip-v2/concept-go-common/exchange/exchangeTypes"
	"github.com/EspressoTrip-v2/concept-go-common/exchange/queue/queueInfo"
	libErrors "github.com/EspressoTrip-v2/concept-go-common/liberrors"
	"github.com/EspressoTrip-v2/concept-go-common/logcodes"
	"github.com/EspressoTrip-v2/concept-go-common/logging"
	"github.com/EspressoTrip-v2/concept-go-common/microservice/microserviceNames"
	"github.com/streadway/amqp"
)

type Consumer interface {
	Listen(key string)
	failOnError(err error, channel *amqp.Channel)
}

type EventConsumer struct {
	rabbitConnection *amqp.Connection
	exchangeName     exchangeNames.ExchangeNames
	exchangeType     exchangeTypes.ExchangeType
	queueName        queueInfo.QueueInfo
	consumerName     string
	serviceName      microserviceNames.MicroserviceNames
	channel          *amqp.Channel
	logger           *logging.LogPublish
}

func NewEventConsumer(rabbitConnection *amqp.Connection, exchangeName exchangeNames.ExchangeNames, exchangeType exchangeTypes.ExchangeType,
	queueName queueInfo.QueueInfo, consumerName string, serviceName microserviceNames.MicroserviceNames) *EventConsumer {
	logger := logging.NewLogPublish(rabbitConnection, serviceName)
	return &EventConsumer{rabbitConnection: rabbitConnection, exchangeName: exchangeName, exchangeType: exchangeType,
		queueName: queueName, consumerName: consumerName, serviceName: serviceName, logger: logger}
}

type ProcessFunc func(msg amqp.Delivery) error

func (c *EventConsumer) Connect(key string) (*EventConsumer, *libErrors.CustomError) {
	var k string
	var err error
	var libError *libErrors.CustomError

	if key != "" {
		k = key
	} else {
		k = string(c.queueName)
	}
	// connect to the rabbit instance
	c.channel, err = c.rabbitConnection.Channel()
	libError = c.failOnError(err)
	if libError != nil {
		return nil, libError
	}

	// declare exchange if not exists
	err = c.channel.ExchangeDeclare(string(c.exchangeName), string(c.exchangeType), true, false, false, false, nil)
	libError = c.failOnError(err)
	if libError != nil {
		return nil, libError
	}
	fmt.Printf("[consumer:%v]: Subscribed on queue:%v\n", c.consumerName, c.queueName)

	// declare a queue if not exists
	q, err := c.channel.QueueDeclare(string(c.queueName), true, false, false, false, nil)
	fmt.Println("Que generated ", q.Name)
	fmt.Println("Que name ", c.queueName)
	libError = c.failOnError(err)
	if libError != nil {
		return nil, libError
	}
	err = c.channel.QueueBind(string(c.queueName), k, string(c.exchangeName), false, nil)
	libError = c.failOnError(err)
	if libError != nil {
		return nil, libError
	}
	return c, nil
}

func (c *EventConsumer) Listen(consumeChannel chan amqp.Delivery) {
	fmt.Println("LIstening")
	deliveredMsg, err := c.channel.Consume(string(c.queueName), "", true, false, false, false, nil)
	if err != nil {
		c.logger.Log(logcodes.ERROR, fmt.Sprintf("go-common library -> Error consumer connecting to queue:%v\n", c.queueName), "events/consumer.go:83", err.Error())
	}
	forever := make(chan bool)
	go func() {
		fmt.Println("GO func run")
		// Receive messages
		for msg := range deliveredMsg {
			fmt.Printf("[consumer:%v]: Message received: %v | queue:%v\n", c.consumerName, c.exchangeName, c.queueName)
			consumeChannel <- msg
			if err != nil {
				c.logger.Log(logcodes.ERROR,
					fmt.Sprintf("go-common library -> Failed process message: %v | queue:%v", c.exchangeName, c.queueName), "events/consumer.go:95", err.Error())
			}
		}
	}()
	<-forever

}

func (c *EventConsumer) failOnError(err error) *libErrors.CustomError {
	if err != nil {
		fmt.Printf("[consumer:%v]: Publisher error: %v | queue:%v | Closing channel\n", c.consumerName, c.exchangeName, c.queueName)
		return libErrors.NewEventConsumerError(err.Error())
	}
	return nil
}
