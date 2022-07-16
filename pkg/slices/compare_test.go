package slices

import (
	. "github.com/mattfenwick/collections/pkg/base"
	"github.com/mattfenwick/collections/pkg/builtins"
	"github.com/mattfenwick/collections/pkg/functions"
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func makePair[A Eq[A], B Eq[B]](a A, b B) *Pair[A, B] {
	return NewPair(a, b)
}

func absoluteValue(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func isPositive[A builtins.Number](a A) bool {
	return a > 0
}

var absoluteValueThenSignKey = OrderedComparatorSplat(
	functions.On(builtins.CompareOrdered[int], absoluteValue),
	functions.On(builtins.CompareBool, isPositive[int]))

var signThenAbsoluteValueKey = OrderedComparatorSplat(
	functions.On(builtins.CompareBool, isPositive[int]),
	functions.On(builtins.CompareOrdered[int], absoluteValue))

func RunCompareTests() {
	Describe("Compare", func() {
		It("slice ordering", func() {
			compare := CompareSlice(builtins.CompareOrdered[int])
			gomega.Expect(compare([]int{}, []int{})).To(gomega.BeEquivalentTo(OrderingEqual))

			gomega.Expect(compare([]int{4}, []int{})).To(gomega.BeEquivalentTo(OrderingGreaterThan))
			gomega.Expect(compare([]int{}, []int{4})).To(gomega.BeEquivalentTo(OrderingLessThan))
			gomega.Expect(compare([]int{4}, []int{4})).To(gomega.BeEquivalentTo(OrderingEqual))

			gomega.Expect(compare([]int{4, 7}, []int{4, 8})).To(gomega.BeEquivalentTo(OrderingLessThan))
			gomega.Expect(compare([]int{4, 7}, []int{4, 6})).To(gomega.BeEquivalentTo(OrderingGreaterThan))
			gomega.Expect(compare([]int{4, 7}, []int{4, 7})).To(gomega.BeEquivalentTo(OrderingEqual))

			gomega.Expect(compare([]int{1, 2, 3}, []int{3, 4, 5})).To(gomega.BeEquivalentTo(OrderingLessThan))
			gomega.Expect(compare([]int{1, 2, 3}, []int{3})).To(gomega.BeEquivalentTo(OrderingLessThan))
			gomega.Expect(compare([]int{1, 2, 3}, []int{})).To(gomega.BeEquivalentTo(OrderingGreaterThan))

			gomega.Expect(compare([]int{3, 4, 5}, []int{1, 2, 3})).To(gomega.BeEquivalentTo(OrderingGreaterThan))
			gomega.Expect(compare([]int{3}, []int{1, 2, 3})).To(gomega.BeEquivalentTo(OrderingGreaterThan))
			gomega.Expect(compare([]int{}, []int{1, 2, 3})).To(gomega.BeEquivalentTo(OrderingLessThan))
		})

		p1 := NewPair[Int, Bool](13, true)
		p2 := NewPair[Int, Bool](14, true)
		p3 := NewPair[Int, Bool](13, false)
		It("functions.On", func() {
			gomega.Expect(functions.On(Compare[Int], First[Int, Bool])(p1, p1)).To(gomega.BeEquivalentTo(OrderingEqual))
			gomega.Expect(functions.On(Compare[Int], First[Int, Bool])(p1, p2)).To(gomega.BeEquivalentTo(OrderingLessThan))
			gomega.Expect(functions.On(Compare[Int], First[Int, Bool])(p1, p3)).To(gomega.BeEquivalentTo(OrderingEqual))
		})

		It("Pair", func() {
			gomega.Expect(ComparePair(p1, p1)).To(gomega.BeEquivalentTo(OrderingEqual))
			gomega.Expect(ComparePair(p1, p2)).To(gomega.BeEquivalentTo(OrderingLessThan))
			gomega.Expect(ComparePair(p1, p3)).To(gomega.BeEquivalentTo(OrderingGreaterThan))
			gomega.Expect(ComparePair(p2, p1)).To(gomega.BeEquivalentTo(OrderingGreaterThan))
			gomega.Expect(ComparePair(p2, p2)).To(gomega.BeEquivalentTo(OrderingEqual))
			gomega.Expect(ComparePair(p2, p3)).To(gomega.BeEquivalentTo(OrderingGreaterThan))
			gomega.Expect(ComparePair(p3, p1)).To(gomega.BeEquivalentTo(OrderingLessThan))
			gomega.Expect(ComparePair(p3, p2)).To(gomega.BeEquivalentTo(OrderingLessThan))
			gomega.Expect(ComparePair(p3, p3)).To(gomega.BeEquivalentTo(OrderingEqual))
		})
	})

	Describe("Eq", func() {
		It("Pair", func() {
			p1 := makePair[Int, Bool](13, true)
			p2 := makePair[Int, Bool](14, true)
			p3 := makePair[Int, Bool](13, false)
			gomega.Expect(EqualPair(p1, p1)).To(gomega.Equal(true))
			gomega.Expect(EqualPair(p1, p2)).To(gomega.Equal(false))
			gomega.Expect(EqualPair(p1, p3)).To(gomega.Equal(false))
			gomega.Expect(EqualPair(p2, p1)).To(gomega.Equal(false))
			gomega.Expect(EqualPair(p2, p2)).To(gomega.Equal(true))
			gomega.Expect(EqualPair(p2, p3)).To(gomega.Equal(false))
			gomega.Expect(EqualPair(p3, p1)).To(gomega.Equal(false))
			gomega.Expect(EqualPair(p3, p2)).To(gomega.Equal(false))
			gomega.Expect(EqualPair(p3, p3)).To(gomega.Equal(true))
		})

		It("slice equality", func() {
			equal := EqualSlice(builtins.Equal[int])
			gomega.Expect(equal([]int{}, []int{})).To(gomega.Equal(true))
			gomega.Expect(equal([]int{}, []int{18, 37})).To(gomega.Equal(false))
			gomega.Expect(equal([]int{25, 39}, []int{})).To(gomega.Equal(false))
			gomega.Expect(equal([]int{14, 32, 65, 8}, []int{14, 32, 65, 8})).To(gomega.Equal(true))
		})
		It("slice equality -- SliceOrd", func() {
			gomega.Expect(Equal[SliceOrd[Int]]([]Int{}, []Int{})).To(gomega.Equal(true))
		})
	})

	Describe("Comparators", func() {
		It("combines correctly", func() {
			gomega.Expect(absoluteValueThenSignKey(-3, 1)).To(gomega.BeEquivalentTo(OrderingGreaterThan))
		})
	})
}
