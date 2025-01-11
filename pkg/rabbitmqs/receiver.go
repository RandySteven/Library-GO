package rabbitmqs_client

import "log"

func (r *RabbitMqClient) Receive(exchange, routingKey string) error {
	err := r.channel.QueueBind(
		r.queue,
		routingKey,
		exchange,
		false,
		nil,
	)
	if err != nil {
		log.Println("failed query bind : ", err)
		return err
	}

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
