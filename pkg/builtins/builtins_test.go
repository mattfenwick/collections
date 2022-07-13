package builtins

import (
	"github.com/mattfenwick/collections/pkg"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func RunBuiltinsTests() {
	Describe("Compare", func() {
		It("Compares strings", func() {
			Expect(Compare[string]("abc", "def")).To(BeEquivalentTo(pkg.OrderingLessThan))
			Expect(Compare[string]("abc", "abc")).To(BeEquivalentTo(pkg.OrderingEqual))
			Expect(Compare[string]("def", "abc")).To(BeEquivalentTo(pkg.OrderingGreaterThan))
		})
	})
}
