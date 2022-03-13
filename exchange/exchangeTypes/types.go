// Package exchangeTypes constant exchange types
package exchangeTypes

type ExchangeType string

const (
	FAN_OUT ExchangeType = "fanout"
	DIRECT  ExchangeType = "direct"
	TOPIC   ExchangeType = "topic"
)
