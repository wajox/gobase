package initializers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/envy"
)

const (
	HTTPServerAddrEnv     = "HTTP_SERVER_ADDR"
	DefaultHTTPServerAddr = ":8000"
)

type HTTPServerAddr string

type HTTPServerConfig struct {
	HTTPServerAddr HTTPServerAddr
	Router         *gin.Engine
}

func InitializeHTTPServerConfig(router *gin.Engine) *HTTPServerConfig {
	return &HTTPServerConfig{
		HTTPServerAddr: HTTPServerAddr(envy.Get(HTTPServerAddrEnv, DefaultHTTPServerAddr)),
		Router:         router,
	}
}

func InitializeHTTPServer(cfg *HTTPServerConfig) (*http.Server, error) {
	// create http server
	srv := &http.Server{
		Addr:    string(cfg.HTTPServerAddr),
		Handler: cfg.Router,
	}

	return srv, nil
}
