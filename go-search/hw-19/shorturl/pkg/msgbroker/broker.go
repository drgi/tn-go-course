package msgbroker

import "context"

type Handler func(context.Context, []byte) error

type MessageBroker interface {
	Send(context.Context, []byte) error
	Register(string, Handler)
}
