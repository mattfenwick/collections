package base

import (
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func makePairOrd[A Ord[A], B Ord[B]](a A, b B) PairOrd[A, B] {
	return PairOrd[A, B](*NewPair(a, b))
}

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

		It("Pair", func() {
			p1 := makePairOrd[Int, Bool](13, true)
			p2 := makePairOrd[Int, Bool](14, true)
			p3 := makePairOrd[Int, Bool](13, false)
			gomega.Expect(p1.Compare(p1)).To(gomega.BeEquivalentTo(OrderingEqual))
			gomega.Expect(p1.Compare(p2)).To(gomega.BeEquivalentTo(OrderingLessThan))
			gomega.Expect(p1.Compare(p3)).To(gomega.BeEquivalentTo(OrderingGreaterThan))
			gomega.Expect(p2.Compare(p1)).To(gomega.BeEquivalentTo(OrderingGreaterThan))
			gomega.Expect(p2.Compare(p2)).To(gomega.BeEquivalentTo(OrderingEqual))
			gomega.Expect(p2.Compare(p3)).To(gomega.BeEquivalentTo(OrderingGreaterThan))
			gomega.Expect(p3.Compare(p1)).To(gomega.BeEquivalentTo(OrderingLessThan))
			gomega.Expect(p3.Compare(p2)).To(gomega.BeEquivalentTo(OrderingLessThan))
			gomega.Expect(p3.Compare(p3)).To(gomega.BeEquivalentTo(OrderingEqual))
		})

		It("SliceOrd", func() {
			gomega.Expect(Compare[SliceOrd[Int]]([]Int{}, []Int{})).To(gomega.BeEquivalentTo(OrderingEqual))

			gomega.Expect(Compare[SliceOrd[Int]]([]Int{4}, []Int{})).To(gomega.BeEquivalentTo(OrderingGreaterThan))
			gomega.Expect(Compare[SliceOrd[Int]]([]Int{}, []Int{4})).To(gomega.BeEquivalentTo(OrderingLessThan))
			gomega.Expect(Compare[SliceOrd[Int]]([]Int{4}, []Int{4})).To(gomega.BeEquivalentTo(OrderingEqual))

			gomega.Expect(Compare[SliceOrd[Int]]([]Int{4, 7}, []Int{4, 8})).To(gomega.BeEquivalentTo(OrderingLessThan))
			gomega.Expect(Compare[SliceOrd[Int]]([]Int{4, 7}, []Int{4, 6})).To(gomega.BeEquivalentTo(OrderingGreaterThan))
			gomega.Expect(Compare[SliceOrd[Int]]([]Int{4, 7}, []Int{4, 7})).To(gomega.BeEquivalentTo(OrderingEqual))

			gomega.Expect(Compare[SliceOrd[Int]]([]Int{1, 2, 3}, []Int{3, 4, 5})).To(gomega.BeEquivalentTo(OrderingLessThan))
			gomega.Expect(Compare[SliceOrd[Int]]([]Int{1, 2, 3}, []Int{3})).To(gomega.BeEquivalentTo(OrderingLessThan))
			gomega.Expect(Compare[SliceOrd[Int]]([]Int{1, 2, 3}, []Int{})).To(gomega.BeEquivalentTo(OrderingGreaterThan))

			gomega.Expect(Compare[SliceOrd[Int]]([]Int{3, 4, 5}, []Int{1, 2, 3})).To(gomega.BeEquivalentTo(OrderingGreaterThan))
			gomega.Expect(Compare[SliceOrd[Int]]([]Int{3}, []Int{1, 2, 3})).To(gomega.BeEquivalentTo(OrderingGreaterThan))
			gomega.Expect(Compare[SliceOrd[Int]]([]Int{}, []Int{1, 2, 3})).To(gomega.BeEquivalentTo(OrderingLessThan))
		})
	})
}
