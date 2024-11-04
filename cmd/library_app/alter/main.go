package main

import (
	"context"
	"fmt"
	"github.com/RandySteven/Library-GO/pkg/configs"
	mysql_client "github.com/RandySteven/Library-GO/pkg/mysql"
	"github.com/RandySteven/Library-GO/queries"
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

	table := ``
	fmt.Print("table : ")
	fmt.Scanln(&table)

	colName := ``
	fmt.Print("column_name : ")
	fmt.Scanln(&colName)

	dataType := ``
	fmt.Print("data_type : ")
	fmt.Scanln(&dataType)

	nullable := ``
	fmt.Print("nullable [y/n] : ")
	fmt.Scanln(&nullable)
	switch nullable {
	case `n`:
		nullable = `DEFAULT ''`
	}

	query := fmt.Sprintf("ALTER TABLE %s ADD %s %s %s", table, colName, dataType, nullable)

	err = cl.Alter(ctx, queries.GoQuery(query))
	if err != nil {
		log.Fatal(err)
		return
	}
}
