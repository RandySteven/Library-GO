package rabbitmqs_client

import (
	"context"
	"github.com/streadway/amqp"
)

type PubSub interface {
	Send(ctx context.Context, exchange, topic string, message any) error
	SendBatch(ctx context.Context, exchange, topic string, messages []any) error
	Receive(ctx context.Context, exchange, routingKey string) (<-chan amqp.Delivery, error)
	ReceiveBatch(ctx context.Context, exchange, routingKey string, maxMessages int) ([]any, error)
	Close() error
	DeclareExchange(ctx context.Context, exchange, typeExchange string) error
}
