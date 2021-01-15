package app_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/wajox/gobase/internal/app"
)

var _ = Describe("Application", func() {
	Describe("InitializeApplication()", func() {
		It("should create new Application", func() {
			app, err := InitializeApplication()

			Expect(app).NotTo(BeNil())
			Expect(err).To(BeNil())
		})
	})
})
