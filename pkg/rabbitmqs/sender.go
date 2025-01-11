package rabbitmqs_client

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func (r *RabbitMqClient) Send(exchange, routingKey string, message interface{}) error {
	err := r.channel.ExchangeDeclare(
		exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		log.Println("failed to exchange declare : ", err)
		return err
	}

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
