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

	Describe("Pair", func() {
		It("ComparePairOrd", func() {
			p1 := NewPair[Int, Bool](13, true)
			p2 := NewPair[Int, Bool](14, true)
			p3 := NewPair[Int, Bool](13, false)
			comparator := ComparePairOrd[Int, Bool]()
			gomega.Expect(comparator(p1, p1)).To(gomega.BeEquivalentTo(OrderingEqual))
			gomega.Expect(comparator(p1, p2)).To(gomega.BeEquivalentTo(OrderingLessThan))
			gomega.Expect(comparator(p1, p3)).To(gomega.BeEquivalentTo(OrderingGreaterThan))
			gomega.Expect(comparator(p2, p1)).To(gomega.BeEquivalentTo(OrderingGreaterThan))
			gomega.Expect(comparator(p2, p2)).To(gomega.BeEquivalentTo(OrderingEqual))
			gomega.Expect(comparator(p2, p3)).To(gomega.BeEquivalentTo(OrderingGreaterThan))
			gomega.Expect(comparator(p3, p1)).To(gomega.BeEquivalentTo(OrderingLessThan))
			gomega.Expect(comparator(p3, p2)).To(gomega.BeEquivalentTo(OrderingLessThan))
			gomega.Expect(comparator(p3, p3)).To(gomega.BeEquivalentTo(OrderingEqual))
		})
		It("ComparePair", func() {
			p1 := NewPair[int, int](13, 71)
			p2 := NewPair[int, int](14, 72)
			p3 := NewPair[int, int](13, 70)
			comparator := ComparePair[int, int]()
			gomega.Expect(comparator(p1, p1)).To(gomega.BeEquivalentTo(OrderingEqual))
			gomega.Expect(comparator(p1, p2)).To(gomega.BeEquivalentTo(OrderingLessThan))
			gomega.Expect(comparator(p1, p3)).To(gomega.BeEquivalentTo(OrderingGreaterThan))
			gomega.Expect(comparator(p2, p1)).To(gomega.BeEquivalentTo(OrderingGreaterThan))
			gomega.Expect(comparator(p2, p2)).To(gomega.BeEquivalentTo(OrderingEqual))
			gomega.Expect(comparator(p2, p3)).To(gomega.BeEquivalentTo(OrderingGreaterThan))
			gomega.Expect(comparator(p3, p1)).To(gomega.BeEquivalentTo(OrderingLessThan))
			gomega.Expect(comparator(p3, p2)).To(gomega.BeEquivalentTo(OrderingLessThan))
			gomega.Expect(comparator(p3, p3)).To(gomega.BeEquivalentTo(OrderingEqual))
		})
	})
}
