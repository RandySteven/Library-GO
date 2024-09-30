package apps

import (
	"context"
	"database/sql"
	"github.com/RandySteven/Library-GO/caches"
	handlers2 "github.com/RandySteven/Library-GO/handlers"
	algolia_client "github.com/RandySteven/Library-GO/pkg/algolia"
	aws_client "github.com/RandySteven/Library-GO/pkg/aws"
	"github.com/RandySteven/Library-GO/pkg/caches"
	"github.com/RandySteven/Library-GO/pkg/configs"
	"github.com/RandySteven/Library-GO/pkg/mysql"
	repositories2 "github.com/RandySteven/Library-GO/repositories"
	schedulers2 "github.com/RandySteven/Library-GO/schedulers"
	usecases2 "github.com/RandySteven/Library-GO/usecases"
	"github.com/go-redis/redis/v8"
)

type App struct {
	AlgoliaSearch *algolia_client.AlgoliaAPISearchClient
	AWSClient     *aws_client.AWSClient
	MySQLDB       *sql.DB
	Redis         *redis.Client
}

func NewApp(config *configs.Config) (*App, error) {
	mysqlDB, err := mysql_client.NewMySQLClient(config)
	if err != nil {
		return nil, err
	}

	err = mysqlDB.Ping()
	if err != nil {
		return nil, err
	}

	redis, err := caches_client.NewRedisCache(config)
	if err != nil {
		return nil, err
	}

	aws, err := aws_client.NewAWSClient(config)
	if err != nil {
		return nil, err
	}

	algolia, err := algolia_client.NewAlgoliaSearch(config)
	if err != nil {
		return nil, err
	}

	return &App{
		MySQLDB:       mysqlDB.Client(),
		Redis:         redis.Client(),
		AWSClient:     aws,
		AlgoliaSearch: algolia,
	}, nil
}

func (app *App) PrepareTheHandler() *handlers2.Handlers {
	repositories := repositories2.NewRepositories(app.MySQLDB)
	caches := caches.NewCaches(app.Redis)
	usecases := usecases2.NewUsecases(repositories, caches, app.AWSClient, app.AlgoliaSearch)
	return handlers2.NewHandlers(usecases)
}

func (app *App) PrepareScheduler() *schedulers2.Schedulers {
	repositories := repositories2.NewRepositories(app.MySQLDB)
	schedulers := schedulers2.NewSchedulers(repositories)
	return schedulers
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
