package main

import (
	"context"
	"github.com/RandySteven/Library-GO/apps"
	"github.com/RandySteven/Library-GO/middlewares"
	"github.com/RandySteven/Library-GO/pkg/configs"
	crons_client "github.com/RandySteven/Library-GO/pkg/crons"
	"github.com/RandySteven/Library-GO/routes"
	"github.com/gorilla/mux"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
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

	schedulers := app.PrepareScheduler()
	scheduler := crons_client.NewScheduler(schedulers)
	if err = scheduler.RunAllJobs(ctx); err != nil {
		log.Fatal(err)
		return
	}

	handlers := app.PrepareTheHandler()
	r := mux.NewRouter()
	routers := routes.NewEndpointRouters(handlers)
	r.Use(middlewares.CorsMiddleware)
	routes.InitRouters(routers, r)
	go config.Run(r)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	//if err = caches.ClearCache(ctx); err != nil {
	//	log.Fatal(err)
	//	return
	//}
}
