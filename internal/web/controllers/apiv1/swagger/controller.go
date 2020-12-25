package swagger

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"                // gin-swagger middleware
	swaggerFiles "github.com/swaggo/gin-swagger/swaggerFiles" // swagger embed files

	//nolint: golint //reason: blank import because of swagger docs init
	_ "github.com/wajox/gobase/api"
	"github.com/wajox/gobase/internal/web/controllers/apiv1"
)

var (
	_ apiv1.Controller = (*Controller)(nil)
)

// Controller implements controller for swagger
type Controller struct {
	apiv1.BaseController
}

// NewController create new instance for swagger controller
func NewController() *Controller {
	return &Controller{}
}

// DefineRoutes adds swagger controller routes to the router
func (ctrl *Controller) DefineRoutes(r gin.IRouter) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
