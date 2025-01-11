package main

import (
	"context"
	"github.com/RandySteven/Library-GO/apps"
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

	if err = app.RabbitMQClient.Receive(); err != nil {
		log.Fatal("consumer error ", err)
		return
	}
	defer app.RabbitMQClient.Close()
}
