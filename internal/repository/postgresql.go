package repository

import (
	"context"
	"fizzbuzz/config"
	"fmt"

	pgx "github.com/jackc/pgx/v5"
	pgconn "github.com/jackc/pgx/v5/pgconn"
)

// pgx interface for DI
type PgxIface interface {
	Begin(context.Context) (pgx.Tx, error)
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	Ping(context.Context) error
	Prepare(context.Context, string, string) (*pgconn.StatementDescription, error)
	Close(context.Context) error
}

const (
	connString = "postgres://%s:%s@localhost:%s/%s?sslmode=disable"
)

// setup DB, initialize pool connection to postgresql
func SetupDB(config config.AppConfig) (*pgx.Conn, error) {
	c, err := pgx.Connect(context.Background(), fmt.Sprintf(connString,
		config.DbUser,
		config.DbPassword,
		config.DbPort,
		config.DbHost))

	if err != nil {
		return nil, err
	}

	return c, nil
}
