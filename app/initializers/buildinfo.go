package initializers

import (
	"github.com/ildarusmanov/gobase/app/build"
)

func InitializeBuildInfo() *build.Info {
	return build.NewInfo()
}
