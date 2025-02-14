package rabbitmqs_client

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func (r *RabbitMqClient) Send(ctx context.Context, exchange, routingKey string, message any) error {
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = r.channel.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return fmt.Errorf("failed to publish message: %w", err)
	}
	log.Println("Message published to exchange:", exchange)
	return nil
}

func (r *RabbitMqClient) SendBatch(ctx context.Context, exchange, routingKey string, messages []any) error {
	for _, message := range messages {
		if err := r.Send(ctx, exchange, routingKey, message); err != nil {
			return err
		}
	}
	return nil
}
