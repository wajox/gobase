package app

import (
	"context"
	"net/http"

	"github.com/wajox/gobase/internal/app/dependencies"
	"github.com/wajox/gobase/internal/app/initializers"
	"github.com/rs/zerolog/log"
)

type Application struct {
	httpServer *http.Server
	Container  *dependencies.Container
}

func InitializeApplication() (*Application, error) {
	if err := initializers.InitializeEnvs(); err != nil {
		return nil, err
	}

	if err := initializers.InitializeLogs(); err != nil {
		return nil, err
	}

	return BuildApplication()
}

func (a *Application) Start(ctx context.Context, cli bool) {
	if cli {
		return
	}

	a.StartHTTPServer()
}

func (a *Application) StartHTTPServer() {
	go func() {
		log.Info().Str("HTTPServerAddress", a.httpServer.Addr).Msg("started http server")

		// service connections
		if err := a.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Panic().Err(err).Msg("HTTP Server stopped")
		}
	}()
}

func (a *Application) Stop() (err error) {
	return a.httpServer.Shutdown(context.TODO())
}
