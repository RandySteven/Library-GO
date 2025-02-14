package consumers

import (
	"context"
	"github.com/RandySteven/Library-GO/apps"
	"github.com/RandySteven/Library-GO/caches"
	consumers_interfaces "github.com/RandySteven/Library-GO/interfaces/consumers"
	"github.com/RandySteven/Library-GO/repositories"
	"log"
	"sync"
)

type (
	Consumers struct {
		BookConsumer consumers_interfaces.BookConsumer
	}

	ConsumerAction func(ctx context.Context) error
)

func NewConsumers(app *apps.App) *Consumers {
	if app == nil || app.RabbitMQClient == nil || app.MySQLDB == nil || app.Redis == nil {
		log.Println("🚨 App dependencies are not properly initialized")
		return nil
	}

	repos := repositories.NewRepositories(app.MySQLDB.Client())
	cache := caches.NewCaches(app.Redis.Client())
	return &Consumers{
		BookConsumer: newBookConsumer(app.RabbitMQClient, repos.BookRepo, cache.BookCache, nil),
	}
}

func (c *Consumers) ListConsumers(ctx context.Context) []error {
	log.Println("book consumer nil gk ya ? ", c.BookConsumer == nil)
	return []error{
		//c.BookConsumer.ConsumeBookToAddElastic(ctx),
		c.BookConsumer.ConsumeBookAddToRedis(ctx),
	}
}

func (c *Consumers) RegisterConsumers(ctx context.Context, errs []error) []error {
	var wg sync.WaitGroup
	var errors []error
	var mu sync.Mutex

	for _, err := range errs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err != nil {
				mu.Lock()
				errors = append(errors, err)
				mu.Unlock()
				log.Printf("Check Consumer error: %v", err)
			}
		}()
	}

	wg.Wait()
	return errors
}

func (c *Consumers) RunConsumers(ctx context.Context) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Start consumers
	consumerActions := c.ListConsumers(ctx)
	errs := c.RegisterConsumers(ctx, consumerActions)

	// Handle errors
	if len(errs) > 0 {
		for _, err := range errs {
			log.Printf("Find Consumer error: %v", err)
		}
		return errs[0] // Return the first error
	}
	<-ctx.Done()
	return nil
}
