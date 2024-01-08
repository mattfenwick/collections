package unstable

import (
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"testing"
)

func TestSuite(t *testing.T) {
	gomega.RegisterFailHandler(Fail)

	RunTableTests()

	RunSpecs(t, "unstable suite")
}
