package mysql

import (
	"database/sql"
	"fmt"
	"github.com/RandySteven/Library-GO/pkg/configs"
)

func NewMySQLClient(config *configs.Config) (*sql.DB, error) {
	mysql := config.Config.MySQL
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", mysql.Host, mysql.Port, mysql.Username, mysql.Password, mysql.Database)
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
