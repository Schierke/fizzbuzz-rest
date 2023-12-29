package cmd

import (
	"fizzbuzz/internal/domain/service"
	"fizzbuzz/internal/handler"
	"fizzbuzz/internal/repository"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start server and listen to HTTP Request",
	RunE: func(_ *cobra.Command, args []string) error {
		// setup DB
		pool, err := repository.SetupDB(connString)
		if err != nil {
			log.Fatal().Msg("Can't setup connection with DB. Abort")
			return err
		}

		// setting up route & middleware
		r := chi.NewRouter()
		r.Use(middleware.Logger)
		r.Use(middleware.Timeout(120 * time.Second))

		// Repository DI
		fizzbuzzRepo := repository.NewFizzBuzzRepository(pool)
		if err != nil {
			return err
		}

		//HTTP Handler
		handler.NewFizzBuzzHandler(service.NewFizzBuzzService(fizzbuzzRepo), r)
		http.ListenAndServe(":8080", r)
		return nil
	},
}
