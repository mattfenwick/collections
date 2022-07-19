package slice

import (
	. "github.com/mattfenwick/collections/pkg/base"
	"github.com/mattfenwick/collections/pkg/function"
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func RunSortTests() {
	Describe("Sort", func() {
		It("Empty", func() {
			gomega.Expect(Sort([]String{})).To(gomega.Equal([]String{}))
		})
		It("Single element", func() {
			gomega.Expect(Sort([]String{"3"})).To(gomega.Equal([]String{"3"}))
		})
		It("Already sorted", func() {
			gomega.Expect(Sort([]Int{-3, 12, 37, 45, 59})).To(gomega.Equal([]Int{-3, 12, 37, 45, 59}))
		})
		It("Duplicate elements", func() {
			gomega.Expect(Sort([]Int{2, 1, 2, 5, 2, 0, 2})).To(gomega.Equal([]Int{0, 1, 2, 2, 2, 2, 5}))
		})
	})
	Describe("SortOn", func() {
		ints := []int{1, 18, -34, 79, 97, 36, 42, -18, -3, -1, -18}
		It("key of self", func() {
			sorted := SortOn(function.Id[Int], Map(WrapInt, ints))
			gomega.Expect(sorted).To(gomega.Equal([]Int{-34, -18, -18, -3, -1, 1, 18, 36, 42, 79, 97}))
		})
		It("key of Pair -- abs, then sign", func() {
			sorted := SortBy(absoluteValueThenSignKey, ints)
			gomega.Expect(sorted).To(gomega.Equal([]int{-1, 1, -3, -18, -18, 18, -34, 36, 42, 79, 97}))
		})
		It("key of Pair -- sign, then abs", func() {
			sorted := SortBy(signThenAbsoluteValueKey, ints)
			gomega.Expect(sorted).To(gomega.Equal([]int{-1, -3, -18, -18, -34, 1, 18, 36, 42, 79, 97}))
		})
	})
	Describe("SortOnBy", func() {
		ints := []Int{1, 18, -34, 79, 97, 36, 42, -18, -3, -1, -18}
		It("key of self", func() {
			sorted := SortOnBy(function.Id[Int], Compare[Int], ints)
			gomega.Expect(sorted).To(gomega.Equal([]Int{-34, -18, -18, -3, -1, 1, 18, 36, 42, 79, 97}))
		})
		// TODO redo these tests
		//It("key of Pair -- abs, then sign", func() {
		//	sorted := SortOnBy(absoluteValueThenSignKey, Compare[PairOrd[Int, Bool]], ints)
		//	gomega.Expect(sorted).To(gomega.Equal([]Int{-1, 1, -3, -18, -18, 18, -34, 36, 42, 79, 97}))
		//})
		//It("key of Pair -- sign, then abc", func() {
		//	sorted := SortOnBy(signThenAbsoluteValueKey, Compare[PairOrd[Bool, Int]], ints)
		//	gomega.Expect(sorted).To(gomega.Equal([]Int{-1, -3, -18, -18, -34, 1, 18, 36, 42, 79, 97}))
		//})
	})
}
