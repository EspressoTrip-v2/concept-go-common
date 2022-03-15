package grpchannel

type GrpcConnectionStatus int

const (
	CLOSE = iota
	OPEN
	ERROR
)
