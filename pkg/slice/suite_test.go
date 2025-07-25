package slice

import (
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"testing"
)

func TestMatcher(t *testing.T) {
	gomega.RegisterFailHandler(Fail)

	RunDataListTests()
	RunSortTests()
	RunCompareTests()
	RunEqualTests()
	RunHelpersTests()

	RunSpecs(t, "slice suite")
}
