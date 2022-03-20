package rabbitmq

import (
	"fmt"
	libErrors "github.com/EspressoTrip-v2/concept-go-common/liberrors"
	"github.com/streadway/amqp"
)

var rabbitClient *RabbitConfig

type RabbitConfig struct {
	connection *amqp.Connection
	channels   map[string]*amqp.Channel
	clientName string
}

func (c *RabbitConfig) channelExists(key string) bool {
	for k, _ := range c.channels {
		if k == key {
			return true
		}
	}
	return false
}

func (c *RabbitConfig) GetChannel(key string) (*amqp.Channel, *libErrors.CustomError) {
	if !c.channelExists(key) {
		return nil, libErrors.NewRabbitConnectionError(fmt.Sprintf("Channel %v has not been added, please add channel first", key))
	}
	return c.channels[key], nil
}

func (c *RabbitConfig) AddChannel(key string) (*amqp.Channel, *libErrors.CustomError) {
	newChannel, err := c.connection.Channel()
	if err != nil {
		return nil, libErrors.NewRabbitConnectionError(fmt.Sprintf("Channel ceration error %v", key))
	}
	c.channels[key] = newChannel
	return newChannel, nil
}

func GetRabbitClient(rabbitUri string, clientName string) (*RabbitConfig, *libErrors.CustomError) {
	var err error
	if rabbitClient == nil {
		rabbitClient = &RabbitConfig{
			connection: nil,
			channels:   make(map[string]*amqp.Channel),
			clientName: clientName,
		}
		rabbitClient.connection, err = amqp.Dial(rabbitUri)
		if err != nil {
			return nil, libErrors.NewRabbitConnectionError(fmt.Sprintf("Rabbit connection error: %v", err.Error()))
		}
		return rabbitClient, nil
	} else {
		return rabbitClient, nil
	}
}
