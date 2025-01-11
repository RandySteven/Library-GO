package rabbitmqs_client

type PubSub interface {
	Send(exchange, topic string, message interface{}) (err error)
	Receive() (err error)
}
