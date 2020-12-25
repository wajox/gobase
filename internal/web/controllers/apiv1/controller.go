package apiv1

import (
	"github.com/gin-gonic/gin"
)

// Controller is an interface for HTTP controllers
type Controller interface {
	DefineRoutes(gin.IRouter)
}
