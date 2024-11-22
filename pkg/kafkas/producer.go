package kafkas_client

import "github.com/IBM/sarama"

func (kc *KafkaClient) SendMessage(topic string, key string, value string) error {
	msg := &sarama.ConsumerMessage{
		Topic: topic,
		Key:   []byte(sarama.StringEncoder(key)),
		Value: []byte(sarama.StringEncoder(value)),
	}
	err := kc.producer.AddMessageToTxn(msg, "", nil)
	if err != nil {
		return err
	}
	return nil
}
