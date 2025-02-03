package mysql_client

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/RandySteven/Library-GO/pkg/configs"
	"github.com/RandySteven/Library-GO/queries"
	"github.com/RandySteven/Library-GO/utils"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

type MySQL interface {
	Close()
	Ping() error
	Client() *sql.DB
	Alter(ctx context.Context, queryScript queries.GoQuery) error
}

type MySQLClient struct {
	db *sql.DB
}

var _ MySQL = &MySQLClient{}

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

func (c *MySQLClient) Alter(ctx context.Context, queryScript queries.GoQuery) error {
	err := utils.QueryValidation(queryScript, `ALTER`)
	if err != nil {
		return err
	}
	isExec := ``
	fmt.Printf("Are u sure want to exec this query %s ? [y/n]:", queryScript.ToString())
	fmt.Scanln(&isExec)

	if isExec == `n` || isExec == `N` {
		return fmt.Errorf(`execution terminate`)
	}

	_, err = c.db.ExecContext(ctx, queryScript.ToString())
	if err != nil {
		return err
	}
	return nil
}
