package rabbitmqs_client

type PubSub interface {
	Send(topic string, message interface{}) (err error)
	Receive(topic string) (err error)
}
