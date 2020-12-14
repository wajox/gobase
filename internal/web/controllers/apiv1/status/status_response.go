package status

import (
	"github.com/ildarusmanov/gobase/internal/app/build"
)

type Response struct {
	ID     string      `jsonapi:"primary,status"`
	Status string      `jsonapi:"attr,status"`
	Build  *build.Info `jsonapi:"attr,build"`
}
