package cmd

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func init() {

}

var rootCmd = &cobra.Command{
	Use:          "fizzbuzz",
	Short:        "fizzbuzz backend entrance",
	SilenceUsage: true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal().Msg("Failed to start")
		os.Exit(1)
	}
}
