package main

import (
	"context"
	"github.com/RandySteven/Library-GO/apps"
	consumers2 "github.com/RandySteven/Library-GO/consumers"
	"github.com/RandySteven/Library-GO/pkg/configs"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load("./files/env/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
}

func main() {
	_ = context.TODO()
	ctx := context.Background()
	configPath, err := configs.ParseFlags()

	if err != nil {
		log.Fatal(err)
		return
	}

	config, err := configs.NewConfig(configPath)
	if err != nil {
		log.Fatal(err)
		return
	}

	app, err := apps.NewApp(config)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = app.Ping()
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("rabbit mq nya nil gk ya ? ", app.RabbitMQClient == nil)
	consumersEs := consumers2.NewConsumers(app)
	log.Println("si consumer nya nil gak ya ? ", consumersEs == nil)
	go func() {
		err = consumersEs.RunConsumers(ctx)
		if err != nil {
			log.Fatal("consumer error run consumers ", err)
		}
	}()

	<-ctx.Done()
	log.Println("application shutdown")

	//if _, err = app.RabbitMQClient.Receive(ctx, "dev_checker", "dev-send-message"); err != nil {
	//	log.Fatal("consumer error ", err)
	//	return
	//}
	//defer app.RabbitMQClient.Close()
}
