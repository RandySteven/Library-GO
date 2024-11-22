package kafkas_client

import "github.com/IBM/sarama"

func (kc *KafkaClient) ConsumeMessages(topic string, partition int32) (<-chan *sarama.ConsumerMessage, error) {
	partitionConsumer, err := kc.consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
	if err != nil {
		return nil, err
	}

	return partitionConsumer.Messages(), nil
}

func (kc *KafkaClient) CloseConsumer() error {
	return kc.consumer.Close()
}
