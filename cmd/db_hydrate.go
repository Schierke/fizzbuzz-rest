package cmd

import (
	"fizzbuzz/config"
	dbhydrate "fizzbuzz/pkg/db_hydrate"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// TODO: move it to config file
func init() {
	rootCmd.AddCommand(hydrateCmd)
}

var hydrateCmd = &cobra.Command{
	Use:   "hydrate",
	Short: "migrating database when starting in a new env",
	RunE: func(_ *cobra.Command, args []string) error {
		appConfig, err := config.LoadAppConfig(".")
		if err != nil {
			log.Fatal().Err(err).Msg("Unable to load configuration")
			return err
		}
		dbhydrate.Start(appConfig)
		return nil
	},
}
