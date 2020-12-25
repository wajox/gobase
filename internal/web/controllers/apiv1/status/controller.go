package status

import (
	"github.com/gin-gonic/gin"
	"github.com/wajox/gobase/internal/app/build"
	"github.com/wajox/gobase/internal/web/controllers/apiv1"
	"github.com/wajox/gobase/internal/web/render"

	"net/http"
)

var (
	_ apiv1.Controller = (*Controller)(nil)
)

// Controller is a controller implementation for status checks
type Controller struct {
	apiv1.BaseController
	buildInfo *build.Info
}

// NewController creates new status controller instance
func NewController(bi *build.Info) *Controller {
	return &Controller{
		buildInfo: bi,
	}
}

// GetStatus godoc
// @Summary Get Application Status
// @Description get status
// @ID get-status
// @Accept json
// @Produce json
// @Success 200 {object} ResponseDoc
// @Router /api/v1/status [get]
func (ctrl *Controller) GetStatus(ctx *gin.Context) {
	render.JSONAPIPayload(ctx, http.StatusOK, &Response{
		Status: http.StatusText(http.StatusOK),
		Build:  ctrl.buildInfo,
	})
}

// DefineRoutes adds controller routes to the router
func (ctrl *Controller) DefineRoutes(r gin.IRouter) {
	r.GET("/api/v1/status", ctrl.GetStatus)
}
