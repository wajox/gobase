package initializers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gobuffalo/envy"
)

const (
	// HTTPServerAddrEnv is an environment variable name for HTTP server address
	HTTPServerAddrEnv = "HTTP_SERVER_ADDR"
	// DefaultHTTPServerAddr  is a default value for HTTP server address
	DefaultHTTPServerAddr = ":8000"
)

// HTTPServerAddr is a type alias for HTTP server address values
type HTTPServerAddr string

// HTTPServerConfig is a type to store HTTP server config
type HTTPServerConfig struct {
	HTTPServerAddr HTTPServerAddr
	Router         *gin.Engine
}

// InitializeHTTPServerConfig initializes new config for InitializeHTTPServer
func InitializeHTTPServerConfig(router *gin.Engine) *HTTPServerConfig {
	return &HTTPServerConfig{
		HTTPServerAddr: HTTPServerAddr(envy.Get(HTTPServerAddrEnv, DefaultHTTPServerAddr)),
		Router:         router,
	}
}

// InitializeHTTPServer create new http.Server instance
func InitializeHTTPServer(cfg *HTTPServerConfig) (*http.Server, error) {
	// create http server
	srv := &http.Server{
		Addr:    string(cfg.HTTPServerAddr),
		Handler: cfg.Router,
	}

	return srv, nil
}
