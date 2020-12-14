package initializers

import (
	"github.com/gin-gonic/gin"
	"github.com/ildarusmanov/gobase/internal/app/dependencies"
	"github.com/ildarusmanov/gobase/internal/web/controllers/apiv1"
	apiv1Status "github.com/ildarusmanov/gobase/internal/web/controllers/apiv1/status"
	apiv1Swagger "github.com/ildarusmanov/gobase/internal/web/controllers/apiv1/swagger"
	"github.com/ildarusmanov/gobase/internal/web/router"
)

func InitializeRouter(container *dependencies.Container) *gin.Engine {
	r := router.NewRouter()

	ctrls := buildControllers(container)

	for i := range ctrls {
		ctrls[i].DefineRoutes(r)
	}

	return r
}

func buildControllers(container *dependencies.Container) []apiv1.Controller {
	return []apiv1.Controller{
		apiv1Status.NewController(container.BuildInfo),
		apiv1Swagger.NewController(),
	}
}
