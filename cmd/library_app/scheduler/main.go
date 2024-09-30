package main

import (
	"context"
	"github.com/RandySteven/Library-GO/apps"
	"github.com/RandySteven/Library-GO/pkg/configs"
	crons_client "github.com/RandySteven/Library-GO/pkg/crons"
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
	log.Println("Starting scheduler...")
	if err = scheduler.RunAllJobs(ctx); err != nil {
		log.Fatalf("Scheduler encountered an error: %v", err)
		return
	}

	log.Println("Scheduler running...")

	// Channel to listen for quit signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Block until a signal is received
	<-quit

	// Gracefully shutdown the scheduler
	log.Println("Shutting down scheduler...")
	cancel() // Cancel context to stop jobs

	err = scheduler.StopAllJobs(ctx)
	if err != nil {
		log.Fatalf("Failed to stop scheduler: %v", err)
		return
	}

	log.Println("Scheduler stopped. Exiting...")
}
