package cmd

import (
	dbhydrate "fizzbuzz/pkg/db_hydrate"

	"github.com/spf13/cobra"
)

// TODO: move it to config file
func init() {
	rootCmd.AddCommand(hydrateCmd)
}

var (
	connString = "postgres://pg:pass@localhost:5432/crud?sslmode=disable"
)

var hydrateCmd = &cobra.Command{
	Use:   "hydrate",
	Short: "migrating database when starting in a new env",
	RunE: func(_ *cobra.Command, args []string) error {
		err := dbhydrate.Start(connString)
		if err != nil {
			panic(err)
		}
		return nil
	},
}
