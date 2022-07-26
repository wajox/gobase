package apiv1_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestApiv1(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Apiv1 Suite")
}
