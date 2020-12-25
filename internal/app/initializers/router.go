package initializers

import (
	"github.com/gin-gonic/gin"
	"github.com/wajox/gobase/internal/app/dependencies"
	"github.com/wajox/gobase/internal/web/controllers/apiv1"
	apiv1Status "github.com/wajox/gobase/internal/web/controllers/apiv1/status"
	apiv1Swagger "github.com/wajox/gobase/internal/web/controllers/apiv1/swagger"
	"github.com/wajox/gobase/internal/web/router"
)

// InitializeRouter initializes new gin router
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
