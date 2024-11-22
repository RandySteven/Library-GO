package kafkas_client

import (
	"github.com/IBM/sarama"
	"github.com/RandySteven/Library-GO/pkg/configs"
)

type KafkaClient struct {
	producer sarama.AsyncProducer
	consumer sarama.Consumer
}

func NewKafkaClient(config *configs.Config) (*KafkaClient, error) {
	client := &KafkaClient{}
	kafkaConf := config.Config.Kafka
	saramaConf := sarama.NewConfig()

	producer, err := sarama.NewAsyncProducer(kafkaConf.Addrs, saramaConf)
	if err != nil {
		return nil, err
	}
	client.producer = producer

	consumer, err := sarama.NewConsumer(kafkaConf.Addrs, saramaConf)
	if err != nil {
		defer client.producer.Close()
		return nil, err
	}
	client.consumer = consumer

	return client, nil
}
