package cli

import (
	"github.com/spf13/cobra"
)

// NewMigrateCmd implements migrate command
func NewMigrateCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "migrate",
		Aliases: []string{"migrate"},
		Short:   "run migrations for database",
		Run: func(cmd *cobra.Command, args []string) {
			// implement here you migration command
		},
	}
}
