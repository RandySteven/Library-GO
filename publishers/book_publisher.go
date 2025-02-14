package publishers

import (
	"context"
	publishers_interface "github.com/RandySteven/Library-GO/interfaces/publishers"
	rabbitmqs_client "github.com/RandySteven/Library-GO/pkg/rabbitmqs"
)

type bookPublisher struct {
	pubsub rabbitmqs_client.PubSub
}

func (b *bookPublisher) PubBookToElastic(ctx context.Context, rID string, bookID uint64) (id string, err error) {
	err = b.pubsub.Send(ctx, "book-id", "books-message", bookID)
	if err != nil {
		return "", nil
	}
	return rID, nil
}

var _ publishers_interface.BookPublisher = &bookPublisher{}

func newBookPublisher(pubsub rabbitmqs_client.PubSub) *bookPublisher {
	return &bookPublisher{
		pubsub: pubsub,
	}
}
