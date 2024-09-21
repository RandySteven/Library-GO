package mysql

import (
	"context"
	"github.com/RandySteven/Library-GO/queries"
	"log"
)

func initTableMigration() []queries.MigrationQuery {
	return []queries.MigrationQuery{
		queries.UserMigration,
		queries.AuthorMigration,
		queries.GenreMigration,
		queries.BookMigration,
		queries.AuthorBookMigration,
		queries.BookGenreMigration,
	}
}

func (c *MySQLClient) Migration(ctx context.Context) {
	migrations := initTableMigration()

	for _, m := range migrations {
		_, err := c.db.ExecContext(ctx, m.ToString())
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}
