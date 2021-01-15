package build_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/wajox/gobase/internal/app/build"
)

var _ = Describe("Info", func() {
	Describe("NewInfo()", func() {
		It("should create new info object", func() {
			info := NewInfo()

			Expect(info).NotTo(BeNil())
		})
	})
})
