//nolint //reason: TODO. disabled just due to developing process. uncomment it when module will be done
package cmd

import (
	// "os"
	// "os/signal"
	// "syscall"
	// "time"

	// "github.com/ildarusmanov/gobase/internal/app"
	// "github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

const (
	MigrationsPath = "file://db/migrations"
)

func NewMigrateCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "migrate",
		Aliases: []string{"migrate"},
		Short:   "run migrations for database",
		Run: func(cmd *cobra.Command, args []string) {
			// log.Info().Msg("Starting")

			// application, err := app.InitializeApplication()

			// if err != nil {
			// 	log.Error().Err(err).Msg("can not initialize application")
			// } else {
			// 	application.Start(true)

			// 	log.Info().Msg("Started")
			// }

			// run migrations

			// log.Error().Err(application.Stop()).Msg("stop application")

			// time.Sleep(time.Second)
			// log.Info().Msg("Finished")
		},
	}
}
