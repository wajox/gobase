package cli

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// ExecuteRootCmd prepares all CLI commands
func ExecuteRootCmd() {
	c := cobra.Command{}

	c.AddCommand(NewServeCmd())
	c.AddCommand(NewMigrateCmd())
	c.AddCommand(NewSeedCmd())

	if err := c.Execute(); err != nil {
		log.Fatal().Err(err)
	}
}
