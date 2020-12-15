package status

import (
	"github.com/gin-gonic/gin"
	"github.com/wajox/gobase/internal/app/build"
	"github.com/wajox/gobase/internal/web/controllers/apiv1"
	"github.com/wajox/gobase/internal/web/render"

	"net/http"
)

const (
	StatusOK = "OK"
)

var (
	_ apiv1.Controller = (*Controller)(nil)
)

type Controller struct {
	apiv1.BaseController
	buildInfo *build.Info
}

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
		Status: StatusOK,
		Build:  ctrl.buildInfo,
	})
}

func (ctrl *Controller) DefineRoutes(r gin.IRouter) {
	r.GET("/api/v1/status", ctrl.GetStatus)
}
