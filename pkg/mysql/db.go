package mysql_client

import (
	"database/sql"
	"fmt"
	"github.com/RandySteven/Library-GO/pkg/configs"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type MySQLClient struct {
	db *sql.DB
}

func NewMySQLClient(config *configs.Config) (*MySQLClient, error) {
	mysql := config.Config.MySQL
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", mysql.Username, mysql.Password, mysql.Host, mysql.Port, mysql.Database)
	log.Println(url)
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}

	connPool := mysql.ConnPool
	db.SetMaxIdleConns(connPool.MaxIdle)
	db.SetMaxOpenConns(connPool.ConnLimit)
	db.SetConnMaxIdleTime(time.Duration(connPool.IdleTime) * time.Second)

	return &MySQLClient{
		db: db,
	}, nil
}

func (c *MySQLClient) Close() {
	c.db.Close()
}

func (c *MySQLClient) Ping() error {
	return c.db.Ping()
}

func (c *MySQLClient) Client() *sql.DB {
	return c.db
}
