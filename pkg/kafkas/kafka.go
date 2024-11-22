package kafkas_client

import "github.com/IBM/sarama"

type KafkaClient struct {
	producer sarama.AsyncProducer
	consumer sarama.Consumer
}
