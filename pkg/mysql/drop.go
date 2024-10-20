package mysql_client

import (
	"context"
	"github.com/RandySteven/Library-GO/queries"
	"log"
)

func initDrop() []queries.DropQuery {
	return []queries.DropQuery{
		queries.DropRatingTable,
		queries.DropBorrowDetailTable,
		queries.DropBorrowTable,
		queries.DropBagTable,
		queries.DropAuthorBookTable,
		queries.DropBookGenreTable,
		queries.DropAuthorTable,
		queries.DropBookTable,
		queries.DropGenreTable,
		queries.DropUserTable,
	}
}

func (c *MySQLClient) Drop(ctx context.Context) {
	drops := initDrop()

	for _, d := range drops {
		_, err := c.db.ExecContext(ctx, d.ToString())
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}
