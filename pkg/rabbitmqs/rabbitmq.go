package rabbitmqs_client

import (
	"encoding/json"
	"fmt"
	"github.com/RandySteven/Library-GO/pkg/configs"
	"github.com/streadway/amqp"
)

type RabbitMqClient struct {
	conn *amqp.Connection
}

func NewRabbitMQClient(configs *configs.Config) (*RabbitMqClient, error) {
	rabbitMQConf := configs.Config.RabbitMQ
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s", rabbitMQConf.Host, rabbitMQConf.Port))
	if err != nil {
		return nil, err
	}
	return &RabbitMqClient{
		conn: conn,
	}, nil
}

func (r *RabbitMqClient) Send(topic string, message interface{}) (err error) {
	channel, err := r.conn.Channel()
	if err != nil {
		return err
	}
	defer channel.Close()

	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return channel.Publish(
		"",
		topic,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}

func (r *RabbitMqClient) Receive(topic string) (err error) {
	return
}

var _ PubSub = &RabbitMqClient{}
