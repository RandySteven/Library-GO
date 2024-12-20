package caches_client

import (
	"context"
	"fmt"
	"github.com/RandySteven/Library-GO/pkg/configs"
	"github.com/dicedb/dicedb-go"
	"log"
)

type DiceDBClient struct {
	client *dicedb.Client
}

func NewDiceDBClient(config *configs.Config) (*DiceDBClient, error) {
	dice := config.Config.Dice
	client := dicedb.NewClient(&dicedb.Options{
		Addr: fmt.Sprintf("%s:%s", dice.Host, dice.Port),
	})
	defer client.Close()

	ctx := context.Background()

	watch := client.WatchConn(ctx)
	defer watch.Close()

	_, err := watch.ZRangeWatch(ctx, "myset", "0", "-1", "REV", "WITHSCORES")
	if err != nil {
		log.Println("err : ", err)
		return nil, err
	}

	ch := watch.Channel()
	for msg := range ch {
		scores := msg.Data.([]dicedb.Z)
		for _, z := range scores {
			fmt.Printf("Member: %s, Score: %f\n", z.Member, z.Score)
		}
	}

	return &DiceDBClient{
		client: client,
	}, nil
}
