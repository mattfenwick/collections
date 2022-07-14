package pkg

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
