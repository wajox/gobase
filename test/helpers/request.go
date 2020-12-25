package helpers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/wajox/gobase/internal/web/controllers/apiv1"
)

// NewTestHTTPRequest is a HTTP request builder for testing purposes
func NewTestHTTPRequest(method, url string, postBody interface{}) *http.Request {
	b, _ := json.Marshal(postBody)

	req := httptest.NewRequest(method, url, bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")

	return req
}

// ServeTestRequest creates HTTP handler for controller
func ServeTestRequest(
	req *http.Request,
	resp http.ResponseWriter,
	r *gin.Engine,
	ctrl apiv1.Controller,
) {
	ctrl.DefineRoutes(r)

	http.Handler(r).ServeHTTP(resp, req)
}
