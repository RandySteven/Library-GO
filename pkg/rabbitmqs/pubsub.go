package rabbitmqs_client

type PubSub interface {
	Send(exchange, topic string, message interface{}) (err error)
	Receive(exchange, routingKey string) (err error)
}
