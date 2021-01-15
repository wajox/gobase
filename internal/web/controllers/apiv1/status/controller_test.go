package status_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	"github.com/wajox/gobase/internal/app/build"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/wajox/gobase/internal/web/controllers/apiv1/status"
)

var _ = Describe("Controller", func() {
	var (
		statusCtrl *Controller
	)

	BeforeEach(func() {
		gin.SetMode(gin.ReleaseMode)

		info := build.NewInfo()
		statusCtrl = NewController(info)
	})

	It("controller should not be nil", func() {
		Expect(statusCtrl).NotTo(BeNil())
	})

	Describe("GetStatus()", func() {
		It("should return status", func() {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			ctx.Request, _ = http.NewRequest("GET", "/api/v1/status", nil)

			statusCtrl.GetStatus(ctx)

			Expect(w.Result().StatusCode).To(Equal(http.StatusOK))
		})
	})
})
