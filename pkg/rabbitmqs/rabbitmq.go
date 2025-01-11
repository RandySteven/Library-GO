package rabbitmqs_client

import (
	"encoding/json"
	"fmt"
	"github.com/RandySteven/Library-GO/pkg/configs"
	"github.com/streadway/amqp"
	"log"
)

type RabbitMqClient struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   string
}

func NewRabbitMQClient(configs *configs.Config) (*RabbitMqClient, error) {
	rabbitMQConf := configs.Config.RabbitMQ
	connectUrlRabbitMq := fmt.Sprintf("amqp://%s:%s", rabbitMQConf.Host, rabbitMQConf.Port)
	log.Println(connectUrlRabbitMq)
	conn, err := amqp.Dial(connectUrlRabbitMq)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	_, err = channel.QueueDeclare(
		rabbitMQConf.Queue,
		true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	return &RabbitMqClient{
		conn:    conn,
		channel: channel,
		queue:   rabbitMQConf.Queue,
	}, nil
}

func (r *RabbitMqClient) Send(exchange, routingKey string, message interface{}) error {
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return r.channel.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}

func (r *RabbitMqClient) Receive() error {
	msgs, err := r.channel.Consume(
		r.queue, "", true, false, false, false, nil)
	if err != nil {
		return err
	}

	forever := make(chan bool)
	go func() {
		for msg := range msgs {
			log.Printf("Received a message: %s", msg.Body)

			if err := processMessage(msg.Body); err != nil {
				log.Printf("Error processing message: %v", err)
			} else {
				log.Println("Message processed successfully")
			}
		}
	}()

	log.Println("Waiting for messages. To exit press CTRL+C")
	<-forever

	return nil
}

func (r *RabbitMqClient) Close() error {
	if err := r.channel.Close(); err != nil {
		return err
	}
	if err := r.conn.Close(); err != nil {
		return err
	}
	return nil
}

func processMessage(body []byte) error {
	log.Printf("Processing message: %s", string(body))
	return nil
}

var _ PubSub = &RabbitMqClient{}
