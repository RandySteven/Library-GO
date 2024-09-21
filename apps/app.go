package apps

import (
	"context"
	"database/sql"
	handlers2 "github.com/RandySteven/Library-GO/handlers"
	"github.com/RandySteven/Library-GO/pkg/caches"
	"github.com/RandySteven/Library-GO/pkg/configs"
	"github.com/RandySteven/Library-GO/pkg/mysql"
	repositories2 "github.com/RandySteven/Library-GO/repositories"
	usecases2 "github.com/RandySteven/Library-GO/usecases"
	"github.com/algolia/algoliasearch-client-go/v4/algolia/search"
	"github.com/aws/aws-sdk-go/service/bedrock"
	"github.com/go-redis/redis/v8"
)

type App struct {
	AlgoliaSearch *search.APIClient
	Bedrock       *bedrock.Bedrock
	MySQLDB       *sql.DB
	Redis         *redis.Client
}

func NewApp(config *configs.Config) (*App, error) {
	mysqlDB, err := mysql.NewMySQLClient(config)
	if err != nil {
		return nil, err
	}

	redis, err := caches.NewRedisCache(config)
	if err != nil {
		return nil, err
	}

	return &App{
		MySQLDB: mysqlDB.Client(),
		Redis:   redis.Client(),
	}, nil
}

func (app *App) PrepareTheHandler() *handlers2.Handlers {
	repositories := repositories2.NewRepositories(app.MySQLDB)
	usecases := usecases2.NewUsecases(repositories)
	return handlers2.NewHandlers(usecases)
}

func (a *App) Ping() error {
	err := a.MySQLDB.Ping()
	if err != nil {
		return err
	}

	err = a.Redis.Ping(context.TODO()).Err()
	if err != nil {
		return err
	}

	return nil
}
