package dependencies

import (
	"github.com/ildarusmanov/gobase/app/build"
)

// Container is a DI container for application
type Container struct {
	BuildInfo *build.Info
}
