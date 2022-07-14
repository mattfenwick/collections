package pkg

import (
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"testing"
)

func TestMatcher(t *testing.T) {
	gomega.RegisterFailHandler(Fail)

	RunEqTests()
	RunOrdTests()
	RunSortTests()

	RunSpecs(t, "builtins suite")
}
