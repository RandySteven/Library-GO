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
	oauth2_client "github.com/RandySteven/Library-GO/pkg/oauth2"
	rabbitmqs_client "github.com/RandySteven/Library-GO/pkg/rabbitmqs"
	repositories2 "github.com/RandySteven/Library-GO/repositories"
	schedulers2 "github.com/RandySteven/Library-GO/schedulers"
	usecases2 "github.com/RandySteven/Library-GO/usecases"
	"log"
)

type App struct {
	AlgoliaSearch  algolia_client.AlgoliaAPISearch
	AWSClient      aws_client.AWS
	MySQLDB        mysql_client.MySQL
	Redis          caches_client.Redis
	Mailer         emails_client.Mail
	ElasticClient  *elastics_client.ElasticClient
	DiceDB         *caches_client.DiceDBClient
	RabbitMQClient rabbitmqs_client.PubSub
	Oauth2         oauth2_client.Oauth2
}

func NewApp(config *configs.Config) (*App, error) {
	mysqlDB, err := mysql_client.NewMySQLClient(config)
	if err != nil {
		log.Println("conn mysql", err)
		return nil, err
	}

	err = mysqlDB.Ping()
	if err != nil {
		log.Println("ping mysql ", err)
		return nil, err
	}

	redis, err := caches_client.NewRedisCache(config)
	if err != nil {
		log.Println("conn redis ", err)
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

	//diceDb, err := caches_client.NewDiceDBClient(config)
	//if err != nil {
	//	return nil, err
	//}

	oauth2, err := oauth2_client.NewOauth2Client(config)
	if err != nil {
		return nil, err
	}

	rabbitMQ, err := rabbitmqs_client.NewRabbitMQClient(config)
	if err != nil {
		log.Fatal("rabbit mq ", err)
		return nil, err
	}

	return &App{
		MySQLDB:       mysqlDB,
		Redis:         redis,
		AWSClient:     aws,
		AlgoliaSearch: algolia,
		Mailer:        mailer,
		ElasticClient: elastic,
		//DiceDB:        diceDb,
		RabbitMQClient: rabbitMQ,
		Oauth2:         oauth2,
	}, nil
}

func (app *App) PrepareTheHandler() *handlers2.Handlers {
	repositories := repositories2.NewRepositories(app.MySQLDB.Client())
	caches := caches.NewCaches(app.Redis.Client())
	usecases := usecases2.NewUsecases(repositories, caches, app.AWSClient, app.AlgoliaSearch, app.RabbitMQClient, app.Oauth2)
	return handlers2.NewHandlers(usecases)
}

func (app *App) PrepareScheduler() *schedulers2.Schedulers {
	repositories := repositories2.NewRepositories(app.MySQLDB.Client())
	caches := caches.NewCaches(app.Redis.Client())
	schedulers := schedulers2.NewSchedulers(repositories, caches, app.AWSClient)
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
