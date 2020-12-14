package initializers

import (
	"github.com/ildarusmanov/gobase/internal/app/build"
)

func InitializeBuildInfo() *build.Info {
	return build.NewInfo()
}
