package initializers

import (
	"github.com/wajox/gobase/internal/app/build"
)

// InitializeBuildInfo creates new build.Info
func InitializeBuildInfo() *build.Info {
	return build.NewInfo()
}
