package base

import (
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func RunOrdTests() {
	Describe("Ord", func() {
		It("Int", func() {
			gomega.Expect(Compare[Int](13, 35)).To(gomega.BeEquivalentTo(OrderingLessThan))
			gomega.Expect(Compare[Int](13, 3)).To(gomega.BeEquivalentTo(OrderingGreaterThan))
			gomega.Expect(Compare[Int](13, 13)).To(gomega.BeEquivalentTo(OrderingEqual))
		})

		It("Bool", func() {
			gomega.Expect(Compare[Bool](true, false)).To(gomega.BeEquivalentTo(OrderingGreaterThan))
			gomega.Expect(Compare[Bool](false, true)).To(gomega.BeEquivalentTo(OrderingLessThan))
			gomega.Expect(Compare[Bool](false, false)).To(gomega.BeEquivalentTo(OrderingEqual))
			gomega.Expect(Compare[Bool](true, true)).To(gomega.BeEquivalentTo(OrderingEqual))
		})
	})

	Describe("constraints.Ordered", func() {
		It("Compares strings", func() {
			gomega.Expect(CompareOrdered[string]("abc", "def")).To(gomega.BeEquivalentTo(OrderingLessThan))
			gomega.Expect(CompareOrdered[string]("abc", "abc")).To(gomega.BeEquivalentTo(OrderingEqual))
			gomega.Expect(CompareOrdered[string]("def", "abc")).To(gomega.BeEquivalentTo(OrderingGreaterThan))
		})
	})
}
