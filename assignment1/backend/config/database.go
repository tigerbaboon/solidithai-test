package config

import (
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
)

type Option struct {
	Dsn      string
	Host     string
	Port     int64
	Database string
	Username string
	Password string
	TimeZone string
}

func Database() *bun.DB {

	db, err := defaultConfig()
	if err != nil {
		panic(err)
	}

	db.AddQueryHook(bundebug.NewQueryHook())

	return db

}

func defaultConfig() (*bun.DB, error) {

	op := Option{
		Host:     conf("DB_HOST", "127.0.0.1"),
		Port:     conf("DB_PORT", int64(5432)),
		Database: conf("DB_DATABASE", "postgres"),
		Username: conf("DB_USER", "postgres"),
		Password: conf("DB_PASSWORD", ""),
		TimeZone: conf("TZ", "Asia/Bangkok"),
	}

	op.Dsn = fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=disable TimeZone=%s", op.Host, op.Port, op.Database, op.Username, op.Password, op.TimeZone)

	config, err := pgx.ParseConfig(op.Dsn)
	if err != nil {
		panic(err)
	}
	sqldb := stdlib.OpenDB(*config)
	db := bun.NewDB(sqldb, pgdialect.New())

	return db, db.Ping()
}
