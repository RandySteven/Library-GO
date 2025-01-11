package rabbitmqs_client

import (
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

var _ PubSub = &RabbitMqClient{}
