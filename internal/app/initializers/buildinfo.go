package initializers

import (
	"github.com/wajox/gobase/internal/app/build"
)

func InitializeBuildInfo() *build.Info {
	return build.NewInfo()
}
