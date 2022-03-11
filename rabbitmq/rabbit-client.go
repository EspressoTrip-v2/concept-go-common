package rabbitmq

import (
	"fmt"
	libErrors "github.com/EspressoTrip-v2/concept-go-common/lib-errors"
	"github.com/streadway/amqp"
)

type Rabbiter interface {
	Connect() (*amqp.Connection, *libErrors.CustomError)
}

type RabbitClient struct {
	rabbitUrl  string
	clientName string
}

func (r RabbitClient) Connect() (*amqp.Connection, *libErrors.CustomError) {
	conn, err := amqp.Dial(r.rabbitUrl)
	if err != nil {
		fmt.Printf("[rabbit:%v]: Failed to connect", r.clientName)

		return nil, libErrors.NewDatabaseError(err.Error())
	}
	defer conn.Close()
	return conn, nil
}

func NewRabbitClient(rabbitUrl string, clientName string) *RabbitClient {
	return &RabbitClient{rabbitUrl: rabbitUrl, clientName: clientName}
}
