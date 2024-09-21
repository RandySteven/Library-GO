package main

import (
	"context"
	"github.com/RandySteven/Library-GO/pkg/configs"
	"github.com/RandySteven/Library-GO/pkg/mysql"
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

	cl, err := mysql.NewMySQLClient(config)
	if err != nil {
		log.Fatal(err)
		return
	}
	cl.Migration(context.TODO())
}
