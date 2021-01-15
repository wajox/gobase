package app_test

import (
	"context"

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

	Describe("methods", func() {
		var (
			app *Application
		)

		BeforeEach(func() {
			app, _ = InitializeApplication()
		})

		Describe("Start(), Stop()", func() {
			It("should start and stop application", func() {
				ctx, cancel := context.WithCancel(context.Background())
				app.Start(ctx, false)

				defer cancel()

				err := app.Stop()

				Expect(err).To(BeNil())
			})
		})
	})
})
