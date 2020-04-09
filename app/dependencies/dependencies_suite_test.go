package dependencies_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDependencies(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Dependencies Suite")
}
