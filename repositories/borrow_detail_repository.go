package repositories

import "database/sql"

type borrowDetailRepository struct {
	db *sql.DB
}

func newBorrowDetailRepository(db *sql.DB) *borrowDetailRepository {
	return &borrowDetailRepository{
		db: db,
	}
}
