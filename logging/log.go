package logging

import (
	"encoding/json"
	"fmt"
	"github.com/EspressoTrip-v2/concept-go-common/exchange/bindkeys"
	"github.com/EspressoTrip-v2/concept-go-common/exchange/exchangeNames"
	"github.com/EspressoTrip-v2/concept-go-common/exchange/exchangeTypes"
	libErrors "github.com/EspressoTrip-v2/concept-go-common/liberrors"
	"github.com/EspressoTrip-v2/concept-go-common/logcodes"
	"github.com/EspressoTrip-v2/concept-go-common/microservice/microserviceNames"
	"github.com/streadway/amqp"
	"time"
)

type Logger interface {
	Publish(data interface{}) *libErrors.CustomError
	Log(errCode logcodes.LogCodes, message string, origin string, details string)
	failOnError(err error) *libErrors.CustomError
}

// LogPublish  is a new publisher for logging, it is best used by embedding it into a local struct
// like below:
//
// 		type localLoggerConfig struct {
//			serviceName microserviceNames.MicroserviceNames
//			logger      *logging.LogPublish
//		}
//
// This way you can serve it like a singleton
type LogPublish struct {
	rabbitChannel *amqp.Channel
	exchangeName  exchangeNames.ExchangeNames
	exchangeType  exchangeTypes.ExchangeType
	bindKey       bindkeys.BindKey
	publisherName string
	serviceName   microserviceNames.MicroserviceNames
}

// NewLogPublish creates a new publisher
func NewLogPublish(rabbitChannel *amqp.Channel, serviceName microserviceNames.MicroserviceNames) *LogPublish {
	return &LogPublish{
		rabbitChannel: rabbitChannel,
		exchangeName:  exchangeNames.LOG,
		exchangeType:  exchangeTypes.DIRECT,
		publisherName: "log-publisher",
		bindKey:       bindkeys.LOG,
		serviceName:   serviceName,
	}
}

func (l LogPublish) Publish(data interface{}) *libErrors.CustomError {
	var err error
	// declare the exchange if it does not exist
	err = l.rabbitChannel.ExchangeDeclare(string(l.exchangeName), string(l.exchangeType), true, false, false, false, nil)
	if err := l.failOnError(err); err != nil {
		return err
	}
	// serialize data
	marshal, err := json.Marshal(data)
	if err := l.failOnError(err); err != nil {
		return err
	}

	err = l.rabbitChannel.Publish(string(l.exchangeName), string(l.bindKey), true, false, amqp.Publishing{
		ContentType:  "text/plain",
		DeliveryMode: amqp.Persistent,
		Body:         marshal,
	})
	if err := l.failOnError(err); err != nil {
		return err
	} else {
		fmt.Printf("[publisher:%v]: Publish:  exchange:%v | route:%v\n", l.publisherName, l.exchangeName, l.bindKey)
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
		fmt.Printf("[publisher:%v]: Publisher error: exchange:%v | route:%v\n", l.publisherName, l.exchangeName, l.bindKey)
		return libErrors.NewEventPublisherError(err.Error())
	}
	return nil
}
