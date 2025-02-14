package rabbitmqs_client

import (
	"context"
	"github.com/streadway/amqp"
	"log"
)

func (r *RabbitMqClient) Receive(ctx context.Context, exchange, routingKey string) (<-chan amqp.Delivery, error) {
	err := r.channel.QueueBind(r.queue, routingKey, exchange, false, nil)
	if err != nil {
		log.Println("Failed to bind queue:", err)
		return nil, err
	}

	msgs, err := r.channel.Consume(r.queue, "", true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}

func (r *RabbitMqClient) ReceiveBatch(ctx context.Context, exchange, routingKey string, maxMessages int) ([]any, error) {
	messages := make([]any, 0, maxMessages)

	msgChan, err := r.Receive(ctx, exchange, routingKey)
	if err != nil {
		return nil, err
	}

	for i := 0; i < maxMessages; i++ {
		select {
		case msg := <-msgChan:
			messages = append(messages, msg.Body)
		case <-ctx.Done():
			return messages, ctx.Err()
		}
	}

	return messages, nil
}
