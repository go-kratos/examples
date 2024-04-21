package data

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	db    *sql.DB
	redis *redis.Client
}

// NewData .
func NewData(db *sql.DB, redisClient *redis.Client) (*Data, error) {
	_, err := db.Exec("create database IF NOT EXISTS `otel_test`")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec("use `otel_test`")
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
create table if not exists otel_test.greeter(
    id int unsigned not null AUTO_INCREMENT primary key,
    hello varchar(255) not null
)
`)
	if err != nil {
		return nil, err
	}

	return &Data{
		db:    db,
		redis: redisClient,
	}, nil
}
