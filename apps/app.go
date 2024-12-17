package apps

import (
	"context"
	"github.com/RandySteven/Library-GO/caches"
	handlers2 "github.com/RandySteven/Library-GO/handlers"
	algolia_client "github.com/RandySteven/Library-GO/pkg/algolia"
	aws_client "github.com/RandySteven/Library-GO/pkg/aws"
	"github.com/RandySteven/Library-GO/pkg/caches"
	"github.com/RandySteven/Library-GO/pkg/configs"
	elastics_client "github.com/RandySteven/Library-GO/pkg/elastics"
	emails_client "github.com/RandySteven/Library-GO/pkg/emails"
	"github.com/RandySteven/Library-GO/pkg/mysql"
	repositories2 "github.com/RandySteven/Library-GO/repositories"
	schedulers2 "github.com/RandySteven/Library-GO/schedulers"
	usecases2 "github.com/RandySteven/Library-GO/usecases"
)

type App struct {
	AlgoliaSearch *algolia_client.AlgoliaAPISearchClient
	AWSClient     *aws_client.AWSClient
	MySQLDB       *mysql_client.MySQLClient
	Redis         *caches_client.RedisClient
	Mailer        *emails_client.Mailer
	ElasticClient *elastics_client.ElasticClient
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

	mailer, err := emails_client.NewMailtrap(config)
	if err != nil {
		return nil, err
	}

	elastic, err := elastics_client.NewElasticClient(config)
	if err != nil {
		return nil, err
	}

	return &App{
		MySQLDB:       mysqlDB,
		Redis:         redis,
		AWSClient:     aws,
		AlgoliaSearch: algolia,
		Mailer:        mailer,
		ElasticClient: elastic,
	}, nil
}

func (app *App) PrepareTheHandler() *handlers2.Handlers {
	repositories := repositories2.NewRepositories(app.MySQLDB.Client())
	caches := caches.NewCaches(app.Redis.Client())
	usecases := usecases2.NewUsecases(repositories, caches, app.AWSClient, app.AlgoliaSearch)
	return handlers2.NewHandlers(usecases)
}

func (app *App) PrepareScheduler() *schedulers2.Schedulers {
	repositories := repositories2.NewRepositories(app.MySQLDB.Client())
	caches := caches.NewCaches(app.Redis.Client())
	schedulers := schedulers2.NewSchedulers(repositories, caches)
	return schedulers
}

func (app *App) RefreshRedis(ctx context.Context) error {
	return app.Redis.ClearCache(ctx)
}

func (a *App) Ping() error {
	err := a.MySQLDB.Ping()
	if err != nil {
		return err
	}

	err = a.Redis.Ping()
	if err != nil {
		return err
	}

	return nil
}
