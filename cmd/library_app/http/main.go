package main

import (
	"github.com/RandySteven/Library-GO/apps"
	"github.com/RandySteven/Library-GO/pkg/configs"
	"github.com/RandySteven/Library-GO/routes"
	"github.com/gorilla/mux"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
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
	handlers := app.PrepareTheHandler()
	r := mux.NewRouter()
	routes.NewEndpointRouters(handlers)
	go config.Run(r)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	//if err = caches.ClearCache(ctx); err != nil {
	//	log.Fatal(err)
	//	return
	//}
}
