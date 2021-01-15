package app

import (
	"context"
	"net/http"

	"github.com/rs/zerolog/log"
	"github.com/wajox/gobase/internal/app/dependencies"
	"github.com/wajox/gobase/internal/app/initializers"
)

// Application is a main struct for the application that contains general information
type Application struct {
	httpServer *http.Server
	Container  *dependencies.Container
}

// InitializeApplication initializes new application
func InitializeApplication() (*Application, error) {
	initializers.InitializeEnvs()

	if err := initializers.InitializeLogs(); err != nil {
		return nil, err
	}

	return BuildApplication()
}

// Start starts application services
func (a *Application) Start(ctx context.Context, cli bool) {
	if cli {
		return
	}

	a.startHTTPServer()
}

// Stop stops application services
func (a *Application) Stop() (err error) {
	return a.httpServer.Shutdown(context.TODO())
}

func (a *Application) startHTTPServer() {
	go func() {
		log.Info().Str("HTTPServerAddress", a.httpServer.Addr).Msg("started http server")

		// service connections
		if err := a.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Panic().Err(err).Msg("HTTP Server stopped")
		}
	}()
}
