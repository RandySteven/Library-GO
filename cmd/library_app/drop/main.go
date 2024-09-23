package main

import (
	"context"
	"github.com/RandySteven/Library-GO/pkg/configs"
	mysql_client "github.com/RandySteven/Library-GO/pkg/mysql"
	"log"
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

	cl, err := mysql_client.NewMySQLClient(config)
	if err != nil {
		log.Fatal(err)
		return
	}

	cl.Drop(context.TODO())
}
