package router_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/wajox/gobase/internal/web/router"
)

var _ = Describe("Router", func() {
	Describe("NewRouter()", func() {
		It("should create new router", func() {
			r := NewRouter()

			Expect(r).NotTo(BeNil())
		})
	})
})
