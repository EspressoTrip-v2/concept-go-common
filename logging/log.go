package logging

import (
	"encoding/json"
	"fmt"
	"github.com/EspressoTrip-v2/concept-go-common/events"
	"github.com/streadway/amqp"
	"time"
)

type LogPublish struct {
	rabbitConnection *amqp.Connection
	exchangeName     events.ExchangeNames
	exchangeType     events.ExchangeType
	queueName        events.QueueInfo
	publisherName    string
	serviceName      events.MicroserviceNames
}

func NewLogPublish(rabbitConnection *amqp.Connection, serviceName events.MicroserviceNames) *LogPublish {
	return &LogPublish{
		rabbitConnection: rabbitConnection,
		exchangeName:     events.LOG,
		exchangeType:     events.DIRECT,
		queueName:        events.LOG_EVENT,
		publisherName:    "log-publisher",
		serviceName:      serviceName,
	}
}

func (l LogPublish) Publish(data interface{}) {

	ch, err := l.rabbitConnection.Channel()
	l.failOnError(err)
	defer ch.Close()

	// declare the exchange if it does not exist
	err = ch.ExchangeDeclare(string(l.exchangeName), string(l.exchangeType), true, false, false, false, nil)
	l.failOnError(err)
	// serialize data
	marshal, err := json.Marshal(data)
	l.failOnError(err)

	err = ch.Publish(string(l.exchangeName), string(l.queueName), true, false, amqp.Publishing{
		ContentType:  "text/plain",
		DeliveryMode: amqp.Persistent,
		Body:         marshal,
	})
	l.failOnError(err)
}

func (l LogPublish) Log(errCode LogCodes, message string, origin string, details string) {
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

func (l *LogPublish) failOnError(err error) {
	if err != nil {
		fmt.Printf("[publisher:%v]: Publisher error: %v | queue:%v", l.publisherName, l.exchangeName, l.queueName)
	}
}
