package mysql_client

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
		queries.BagMigration,
		queries.BorrowMigration,
		queries.BorrowDetailMigration,
		queries.StoryGeneratorMigration,
		queries.RatingMigration,
		queries.CommentMigration,
		queries.EventMigration,
		queries.EventUserMigration,
		queries.RoomChatMigration,
		//queries.RoomChatUserMigration,
		//queries.ChatMigration,
		queries.RoleMigration,
		queries.RoleUserMigration,
		queries.UserGenreMigration,
		queries.RoomMigration,
		queries.RoomPhotoMigration,
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
