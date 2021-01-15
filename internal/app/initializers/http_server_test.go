package initializers_test

import (
	"github.com/gin-gonic/gin"

	"github.com/wajox/gobase/internal/web/router"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/wajox/gobase/internal/app/initializers"
)

var _ = Describe("HttpServer", func() {
	Describe("InitializeHTTPServerConfig()", func() {
		var (
			r *gin.Engine
		)

		BeforeEach(func() {
			r = router.NewRouter()
		})

		It("should initialize config for http server initializer", func() {
			cfg := InitializeHTTPServerConfig(r)

			Expect(cfg).NotTo(BeNil())
			Expect(cfg.Router).To(Equal(r))
			Expect(cfg.HTTPServerAddr).NotTo(BeEmpty())
		})
	})

	Describe("InitializeHTTPServer()", func() {
		var (
			r   *gin.Engine
			cfg *HTTPServerConfig
		)

		BeforeEach(func() {
			r = router.NewRouter()
			cfg = InitializeHTTPServerConfig(r)
		})

		It("should initialize HTTP server", func() {
			srv, err := InitializeHTTPServer(cfg)

			Expect(srv).NotTo(BeNil())
			Expect(err).To(BeNil())
		})
	})
})
