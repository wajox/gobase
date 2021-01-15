package initializers

import (
	"github.com/gobuffalo/envy"
	"github.com/rs/zerolog/log"
)

// InitializeEnvs intializes envy
func InitializeEnvs() {
	if err := envy.Load(); err != nil {
		log.Info().Err(err).Msg("can not load .env file")

		envy.Reload()
	}
}
