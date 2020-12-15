//+build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/wajox/gobase/internal/app/dependencies"
	"github.com/wajox/gobase/internal/app/initializers"
)

func BuildApplication() (*Application, error) {
	wire.Build(
		initializers.InitializeBuildInfo,
		wire.Struct(new(dependencies.Container), "BuildInfo"),
		initializers.InitializeRouter,
		initializers.InitializeHTTPServerConfig,
		initializers.InitializeHTTPServer,
		wire.Struct(new(Application), "HTTPServer", "Container"),
	)

	return &Application{}, nil
}
