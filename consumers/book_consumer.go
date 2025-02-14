package consumers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/RandySteven/Library-GO/entities/indexes"
	"github.com/RandySteven/Library-GO/entities/models"
	"github.com/RandySteven/Library-GO/entities/payloads/responses"
	caches_interfaces "github.com/RandySteven/Library-GO/interfaces/caches"
	consumers_interfaces "github.com/RandySteven/Library-GO/interfaces/consumers"
	repositories_interfaces "github.com/RandySteven/Library-GO/interfaces/repositories"
	searchers_interfaces "github.com/RandySteven/Library-GO/interfaces/searchers"
	rabbitmqs_client "github.com/RandySteven/Library-GO/pkg/rabbitmqs"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type bookConsumer struct {
	pubsub      rabbitmqs_client.PubSub
	bookIndexer searchers_interfaces.BookSearcher
	bookRepo    repositories_interfaces.BookRepository
	bookCache   caches_interfaces.BookCache
}

func (b *bookConsumer) ConsumeBookToAddElastic(ctx context.Context) error {
	if b.pubsub == nil {
		return fmt.Errorf(`failed to run book to add elastic due nil`)
	}
	msgs, err := b.pubsub.Receive(ctx, "book_exchange", "book-send-message")
	if err != nil {
		return err
	}

	for msg := range msgs {
		log.Println("Received message:", string(msg.Body))

		var bookEntity models.Book
		err := json.Unmarshal(msg.Body, &bookEntity)
		if err != nil {
			log.Println("Error decoding message:", err)
			continue // Skip to the next message if decoding fails
		}

		_, err = b.bookIndexer.SaveSearch(ctx, &indexes.BookIndex{
			ID:          bookEntity.ID,
			Title:       bookEntity.Title,
			Description: bookEntity.Description,
			Image:       bookEntity.Image,
			Status:      bookEntity.Status.ToString(),
		})
		if err != nil {
			log.Println("Error saving to Elasticsearch:", err)
			continue
		}

		if err := msg.Ack(false); err != nil {
			log.Println("Failed to acknowledge message:", err)
		}
	}

	return nil
}

func (b *bookConsumer) ConsumeBookAddToRedis(ctx context.Context) (err error) {
	msgs, err := b.pubsub.Receive(ctx, "book_exchange", "book-send-message")
	if err != nil {
		return err
	}

	for msg := range msgs {
		log.Println("Received message:", string(msg.Body))

		var bookEntity models.Book
		err := json.Unmarshal(msg.Body, &bookEntity)
		if err != nil {
			log.Println("Error decoding message:", err)
			continue
		}
		bookList := &responses.ListBooksResponse{
			ID:        bookEntity.ID,
			Image:     bookEntity.Image,
			Title:     bookEntity.Title,
			Status:    bookEntity.Status.ToString(),
			Rating:    0,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: nil,
		}
		listBookResponse, err := b.bookCache.GetMultiData(ctx)
		if err != nil {
			if !errors.Is(err, redis.Nil) {
				return err
			}
		}
		listBookResponse = append(listBookResponse, bookList)
		err = b.bookCache.SetMultiData(ctx, listBookResponse)
		if err != nil {
			return err
		}
		if err := msg.Ack(false); err != nil {
			log.Println("Failed to acknowledge message:", err)
		}
	}
	log.Println("success to add into redis")
	return nil
}

var _ consumers_interfaces.BookConsumer = &bookConsumer{}

func newBookConsumer(pubsub rabbitmqs_client.PubSub,
	bookRepo repositories_interfaces.BookRepository,
	bookCache caches_interfaces.BookCache,
	bookIndexer searchers_interfaces.BookSearcher) *bookConsumer {

	//err := pubsub.DeclareExchange(context.Background(), "book_exchange", "topic")
	//if err != nil {
	//	return nil
	//}
	return &bookConsumer{
		pubsub:      pubsub,
		bookRepo:    bookRepo,
		bookCache:   bookCache,
		bookIndexer: bookIndexer, // Add this
	}
}
