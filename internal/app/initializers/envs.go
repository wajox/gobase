package initializers

import (
	"github.com/gobuffalo/envy"
	"github.com/rs/zerolog/log"
)

func InitializeEnvs() error {
	if err := envy.Load(); err != nil {
		log.Info().Err(err).Msg("can not load .env file")
	}

	return nil
}
