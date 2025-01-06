package main

import (
	"context"
	"github.com/RandySteven/Library-GO/pkg/configs"
	mysql_client "github.com/RandySteven/Library-GO/pkg/mysql"
	"log"
)

func main() {
	configPath, err := configs.ParseFlags()
	ctx := context.Background()
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

	//uncomment these lines for seed user
	//err = cl.SeedUserData(ctx)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}

	//uncomment these lines for seed genre
	//err = cl.SeedGenreData(ctx)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}

	//uncomment these lines for seed author
	//err = cl.SeedAuthorData(ctx)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}

	//uncomment these lines for seed book
	//err = cl.SeedBookData(ctx)
	//if err != nil {
	//	log.Fatal(err)
	//	return
	//}

	//uncomment these lines for seed role
	err = cl.SeedRoleData(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}

}
