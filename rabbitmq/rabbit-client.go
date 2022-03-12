package rabbitmq

import (
	libErrors "github.com/EspressoTrip-v2/concept-go-common/lib-errors"
	"github.com/streadway/amqp"
)

var rabbitClient *rabbitConfig

type rabbitConfig struct {
	connection *amqp.Connection
	clientName string
}

func StartRabbitClient(rabbitUri string, clientName string) (*amqp.Connection, *libErrors.CustomError) {
	var conn *amqp.Connection
	var err error
	if rabbitClient == nil {
		rabbitClient = &rabbitConfig{
			connection: nil,
			clientName: clientName,
		}
		if conn, err = amqp.Dial(rabbitUri); err != nil {
			return nil, libErrors.NewRabbitConnectionError("Rabbit connection error")
		}
		rabbitClient.connection = conn
	}
	return rabbitClient.connection, nil
}

func IsConnected() bool {
	return !rabbitClient.connection.IsClosed()
}

func GetRabbitConnection() (*amqp.Connection, *libErrors.CustomError) {
	if rabbitClient == nil {
		return nil, libErrors.NewRabbitConnectionError("Rabbit connection does not exist")
	}
	return rabbitClient.connection, nil
}
