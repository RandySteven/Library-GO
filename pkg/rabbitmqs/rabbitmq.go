package rabbitmqs_client

import (
	"context"
	"fmt"
	"github.com/RandySteven/Library-GO/pkg/configs"
	"github.com/streadway/amqp"
	"log"
)

type RabbitMqClient struct {
	conn     *amqp.Connection
	channel  *amqp.Channel
	exchange string
	queue    string
}

func NewRabbitMQClient(configs *configs.Config) (*RabbitMqClient, error) {
	rabbitMQConf := configs.Config.RabbitMQ
	//fmt.Sprintf("amqp://%s:%s@%s:%s/", rabbitMQConf.User, rabbitMQConf.Password, rabbitMQConf.Host, rabbitMQConf.Port)

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

	q, err := channel.QueueDeclare(
		rabbitMQConf.Queue,
		true, false, false, false, nil)
	if err != nil {
		log.Println("failed to declare queue : ", err)
		return nil, err
	}

	return &RabbitMqClient{
		conn:    conn,
		channel: channel,
		queue:   q.Name,
	}, nil
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

func (r *RabbitMqClient) DeclareExchange(ctx context.Context, exchange, typeExchange string) error {
	err := r.channel.ExchangeDeclare(
		exchange,
		typeExchange,
		true,
		false,
		false,
		false,
		nil)
	if err != nil {
		return err
	}
	return nil
}

var _ PubSub = &RabbitMqClient{}
