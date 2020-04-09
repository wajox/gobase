package status

import (
	"github.com/ildarusmanov/gobase/app/build"
	"github.com/ildarusmanov/gobase/app/config"
)

type Response struct {
	ID     string          `jsonapi:"primary,status"`
	Status string          `jsonapi:"attr,status"`
	App    *config.AppInfo `jsonapi:"attr,app"`
	Build  *build.Info     `jsonapi:"attr,build"`
}
