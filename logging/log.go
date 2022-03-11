package logging

import (
	"encoding/json"
	"fmt"
	"github.com/EspressoTrip-v2/concept-go-common/exchange/exchangeNames"
	"github.com/EspressoTrip-v2/concept-go-common/exchange/exchangeTypes"
	"github.com/EspressoTrip-v2/concept-go-common/exchange/queue/queueInfo"
	libErrors "github.com/EspressoTrip-v2/concept-go-common/lib-errors"
	"github.com/EspressoTrip-v2/concept-go-common/logcodes"
	"github.com/EspressoTrip-v2/concept-go-common/microservice/microserviceNames"
	"github.com/streadway/amqp"
	"time"
)

type LogPublish struct {
	rabbitConnection *amqp.Connection
	exchangeName     exchangeNames.ExchangeNames
	exchangeType     exchangeTypes.ExchangeType
	queueName        queueInfo.QueueInfo
	publisherName    string
	serviceName      microserviceNames.MicroserviceNames
}

func NewLogPublish(rabbitConnection *amqp.Connection, serviceName microserviceNames.MicroserviceNames) *LogPublish {
	return &LogPublish{
		rabbitConnection: rabbitConnection,
		exchangeName:     exchangeNames.LOG,
		exchangeType:     exchangeTypes.DIRECT,
		queueName:        queueInfo.LOG_EVENT,
		publisherName:    "log-publisher",
		serviceName:      serviceName,
	}
}

func (l LogPublish) Publish(data interface{}) *libErrors.CustomError {
	ch, err := l.rabbitConnection.Channel()
	if err := l.failOnError(err); err != nil {
		return err
	}
	defer ch.Close()

	// declare the exchange if it does not exist
	err = ch.ExchangeDeclare(string(l.exchangeName), string(l.exchangeType), true, false, false, false, nil)
	if err := l.failOnError(err); err != nil {
		return err
	}
	// serialize data
	marshal, err := json.Marshal(data)
	if err := l.failOnError(err); err != nil {
		return err
	}

	err = ch.Publish(string(l.exchangeName), string(l.queueName), true, false, amqp.Publishing{
		ContentType:  "text/plain",
		DeliveryMode: amqp.Persistent,
		Body:         marshal,
	})
	if err := l.failOnError(err); err != nil {
		return err
	}
	return nil
}

func (l LogPublish) Log(errCode logcodes.LogCodes, message string, origin string, details string) {
	msg := LogMsg{
		Service:    string(l.serviceName),
		LogContext: string(errCode),
		Origin:     origin,
		Message:    message,
		Details:    details,
		Date:       fmt.Sprint(time.Now().Format(time.RFC3339)),
	}
	l.Publish(msg)

}

func (l *LogPublish) failOnError(err error) *libErrors.CustomError {
	if err != nil {
		fmt.Printf("[publisher:%v]: Publisher error: %v | queue:%v\n", l.publisherName, l.exchangeName, l.queueName)
		return libErrors.NewEventPublisherError(err.Error())
	}
	return nil
}
