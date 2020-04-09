package initializers

import (
	"github.com/ildarusmanov/gobase/app/build"
)

func InitializeBuildInfo() (*build.Info, error) {
	return build.NewInfo()
}
