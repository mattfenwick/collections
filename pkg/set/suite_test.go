package set

import (
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"testing"
)

func TestSet(t *testing.T) {
	gomega.RegisterFailHandler(Fail)

	RunSetTests()

	RunSpecs(t, "set suite")
}
