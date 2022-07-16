package builtins

import (
	"github.com/mattfenwick/collections/pkg/base"
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func RunCompareTests() {
	Describe("Compare", func() {
		It("Compares strings", func() {
			gomega.Expect(CompareOrdered[string]("abc", "def")).To(gomega.BeEquivalentTo(base.OrderingLessThan))
			gomega.Expect(CompareOrdered[string]("abc", "abc")).To(gomega.BeEquivalentTo(base.OrderingEqual))
			gomega.Expect(CompareOrdered[string]("def", "abc")).To(gomega.BeEquivalentTo(base.OrderingGreaterThan))
		})
	})
}
