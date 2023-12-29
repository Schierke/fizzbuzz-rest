package dbhydrate

import (
	"context"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5"
)

func Start(pgString string) error {
	// migrating and getting ready for DB.
	c, err := pgx.Connect(context.Background(), pgString)
	if err != nil {
		return err
	}

	// TODO: should also stock somewhere
	tables := []string{
		"user",
	}
	for _, table := range tables {
		dropTable(c, table)
	}

	err = c.Close(context.Background())
	if err != nil {
		return err
	}

	// Re-migrate
	m, err := migrate.New(
		"file://db/migrations",
		pgString)
	if err != nil {
		return err
	}
	return m.Up()
}

func dropTable(c *pgx.Conn, table string) {
	stmnt := fmt.Sprintf(`DROP TABLE IF EXISTS "%v";`, table)
	_, err := c.Exec(context.Background(), stmnt)
	if err != nil {
		fmt.Printf("error dropping database: %v\n", err)
	}
}
