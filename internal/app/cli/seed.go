package cli

import (
	"github.com/spf13/cobra"
)

// NewSeedCmd implements seed command for DB setup
func NewSeedCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "seed",
		Aliases: []string{"seed"},
		Short:   "apply seeds for database",
		Run: func(cmd *cobra.Command, args []string) {
			// implement your seed command for DB setup
		},
	}
}
