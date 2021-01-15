package initializers_test

import (
	"os"

	"github.com/gobuffalo/envy"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/wajox/gobase/internal/app/initializers"
)

var _ = Describe("Envs", func() {
	Describe("InitializeEnvs()", func() {
		var (
			k, v string
		)

		BeforeEach(func() {
			k = "SOME_TEST_ENV"
			v = "SOME_TEST_ENV_VALUE"

			Expect(os.Setenv(k, v)).To(BeNil())
		})

		It("should initialize envs with Envy package", func() {
			InitializeEnvs()

			Expect(os.Getenv(k)).To(Equal(v))
			Expect(envy.Get(k, "")).To(Equal(v))
		})
	})
})
