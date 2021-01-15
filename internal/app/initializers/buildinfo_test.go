package initializers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/wajox/gobase/internal/app/initializers"
)

var _ = Describe("Buildinfo", func() {
	Describe("InitializeBuildInfo", func() {
		It("should initialize and return build.Info", func() {
			info := InitializeBuildInfo()

			Expect(info).NotTo(BeNil())
		})
	})
})
