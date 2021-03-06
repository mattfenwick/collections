package dict

import (
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"testing"
)

func TestMatcher(t *testing.T) {
	gomega.RegisterFailHandler(Fail)

	RunSortTests()
	RunCompareTests()
	RunEqualTests()

	RunSpecs(t, "builtins suite")
}
