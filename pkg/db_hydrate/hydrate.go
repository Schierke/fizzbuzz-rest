package dbhydrate

import (
	"errors"
	"fizzbuzz/config"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/rs/zerolog/log"
)

const (
	connString     = "postgres://%s:%s@%s:%s/%s?sslmode=disable"
	defaultAttempt = 20
	defautTimeout  = time.Second
)

func Start(config config.AppConfig) {
	dbUrl := fmt.Sprintf(connString,
		config.DbUser,
		config.DbPassword,
		config.DbHost,
		config.DbPort,
		config.DbName)

	var (
		attempts = defaultAttempt
		m        *migrate.Migrate
		err      error
	)

	for attempts > 0 {
		m, err = migrate.New(
			"file://db/migrations",
			dbUrl)
		if err == nil {
			break
		}

		log.Info().Msg(fmt.Sprintf("Migrate: postgres is trying to connect, attempts left: %d", attempts))
		time.Sleep(defautTimeout)
		attempts--
	}

	if err != nil {
		log.Fatal().Msg(fmt.Sprintf("Migrate: postgres connect error: %s", err))
	}
	err = m.Down()
	if errors.Is(err, migrate.ErrNoChange) {
		log.Info().Msg("No change on migration")
	} else if err != nil {
		log.Fatal().Msg(fmt.Sprintf("Migrate up error:%s", err))
	}

	err = m.Up()
	defer m.Close()
	if errors.Is(err, migrate.ErrNoChange) {
		log.Info().Msg("No change on migration")
	} else if err != nil {
		log.Fatal().Msg(fmt.Sprintf("Migrate up error:%s", err))
	}
}
