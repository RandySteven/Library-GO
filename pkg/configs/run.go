package configs

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func ParseFlags() (string, error) {
	var configPath string
	flag.StringVar(&configPath, "config", "./config.yml", "path to config file")
	flag.Parse()
	if err := ValidateConfigPath(configPath); err != nil {
		return "", err
	}
	return configPath, nil
}

func ValidateConfigPath(path string) error {
	s, err := os.Stat(path)
	if err != nil {
		return err
	}
	if s.IsDir() {
		return fmt.Errorf("invalid file format")
	}
	return nil
}

func (c *Config) Run(r *mux.Router) {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	corsHandler := cors.AllowAll().Handler(r)

	server := c.Config.Server

	srv := &http.Server{
		Addr:         server.Host + ":" + server.Port,
		Handler:      corsHandler,
		ReadTimeout:  server.Timeout.Read * time.Second,
		WriteTimeout: server.Timeout.Write * time.Second,
		IdleTimeout:  server.Timeout.Idle * time.Second,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	stop()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")
}
