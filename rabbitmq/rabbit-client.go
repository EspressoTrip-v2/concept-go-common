package rabbitmq

import (
	"fmt"
	libErrors "github.com/EspressoTrip-v2/concept-go-common/lib-errors"
	"github.com/streadway/amqp"
)

type RabbitClient struct {
	rabbitUrl  string
	clientName string
}

func (r RabbitClient) Connect() (*amqp.Connection, *libErrors.CustomError) {
	conn, err := amqp.Dial(r.rabbitUrl)
	r.failOnError(err, conn)

	fmt.Printf("[%v:rabbitmq]: Connected successfully\n", r.clientName)
	return conn, nil
}

func (r *RabbitClient) failOnError(err error, conn *amqp.Connection) *libErrors.CustomError {
	if err != nil {
		fmt.Printf("[rabbit:%v]: Connection error: %v\n", r.clientName, err.Error())
		conn.Close()
		return libErrors.NewEventPublisherError(err.Error())
	}
	return nil
}

func NewRabbitClient(rabbitUrl string, clientName string) *RabbitClient {
	return &RabbitClient{rabbitUrl: rabbitUrl, clientName: clientName}
}
