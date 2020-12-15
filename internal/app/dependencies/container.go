package dependencies

import (
	"github.com/wajox/gobase/internal/app/build"
)

// Container is a DI container for application
type Container struct {
	BuildInfo *build.Info
}
