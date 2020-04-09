package status

import (
	"github.com/gin-gonic/gin"
	"github.com/ildarusmanov/gobase/app/config"
	"github.com/ildarusmanov/gobase/web/controllers/apiv1"
	"github.com/ildarusmanov/gobase/web/render"

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
	AppConfig *config.Config
}

func NewController(appConfig *config.Config) *Controller {
	return &Controller{
		AppConfig: appConfig,
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
		App:    ctrl.AppConfig.AppInfo,
		Build:  ctrl.AppConfig.BuildInfo,
	})
}

func (ctrl *Controller) DefineRoutes(r gin.IRouter) {
	r.GET("/api/v1/status", ctrl.GetStatus)
}
