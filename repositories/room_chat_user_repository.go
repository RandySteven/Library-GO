package repositories

import "database/sql"

type roomChatUserRepository struct {
	db *sql.DB
	tx *sql.Tx
}

func newRoomChatUserRepository(db *sql.DB) *roomChatUserRepository {
	return &roomChatUserRepository{
		db: db,
	}
}
