package main

import (
	"context"
	"github.com/RandySteven/Library-GO/apps"
	"github.com/RandySteven/Library-GO/pkg/configs"
	"github.com/RandySteven/Library-GO/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	err := godotenv.Load("./files/env/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}
}

func main() {
	ctx := context.TODO()
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

	handlers := app.PrepareTheHandler()
	r := mux.NewRouter()
	routers := routes.NewEndpointRouters(handlers)
	routes.InitRouters(routers, r)
	go config.Run(r)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	if err = app.RefreshRedis(ctx); err != nil {
		log.Fatal(err)
		return
	}
}
