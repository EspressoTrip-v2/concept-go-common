package rabbitmq

import (
	libErrors "github.com/EspressoTrip-v2/concept-go-common/liberrors"
	"github.com/streadway/amqp"
)

var rabbitClient *RabbitConfig

type RabbitConfig struct {
	connection *amqp.Connection
	clientName string
}

// StartRabbitClient connects to the message bus and returns the pointer to the connection
func StartRabbitClient(rabbitUri string, clientName string) (*RabbitConfig, *libErrors.CustomError) {
	var conn *amqp.Connection
	var err error
	if rabbitClient == nil {
		rabbitClient = &RabbitConfig{
			connection: nil,
			clientName: clientName,
		}
		if conn, err = amqp.Dial(rabbitUri); err != nil {
			return nil, libErrors.NewRabbitConnectionError("Rabbit connection error")
		}
		rabbitClient.connection = conn
	}
	return rabbitClient, nil
}

func (c RabbitConfig) GetConnection() *amqp.Connection {
	return c.connection
}


func IsConnected() bool {
	return !rabbitClient.connection.IsClosed()
}

func GetRabbitConnection() (*RabbitConfig, *libErrors.CustomError) {
	if rabbitClient == nil {
		return nil, libErrors.NewRabbitConnectionError("Rabbit connection does not exist")
	}
	return rabbitClient, nil
}
