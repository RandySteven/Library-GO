package repositories

import "database/sql"

type roleRepository struct {
	db *sql.DB
	tx *sql.Tx
}
