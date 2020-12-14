package apiv1

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	DefineRoutes(gin.IRouter)
}
