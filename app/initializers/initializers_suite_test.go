package initializers_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestInitializers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Initializers Suite")
}
