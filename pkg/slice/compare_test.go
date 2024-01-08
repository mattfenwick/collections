package slice

import (
	. "github.com/mattfenwick/collections/pkg/base"
	"github.com/mattfenwick/collections/pkg/builtin"
	"github.com/mattfenwick/collections/pkg/function"
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func absoluteValue(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func isPositive[A builtin.Number](a A) bool {
	return a > 0
}

var absoluteValueThenSignKey = CompareBy(
	function.On(CompareOrdered[int], absoluteValue),
	function.On(CompareBool, isPositive[int]))

var signThenAbsoluteValueKey = CompareBy(
	function.On(CompareBool, isPositive[int]),
	function.On(CompareOrdered[int], absoluteValue))

func RunCompareTests() {
	Describe("Compare", func() {
		It("slice ordering", func() {
			compare := CompareSlicePairwiseBy(CompareOrdered[int])
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
			gomega.Expect(function.On(Compare[Int], Fst[Int, Bool])(p1, p1)).To(gomega.BeEquivalentTo(OrderingEqual))
			gomega.Expect(function.On(Compare[Int], Fst[Int, Bool])(p1, p2)).To(gomega.BeEquivalentTo(OrderingLessThan))
			gomega.Expect(function.On(Compare[Int], Fst[Int, Bool])(p1, p3)).To(gomega.BeEquivalentTo(OrderingEqual))
		})
	})

	Describe("Comparators", func() {
		It("combines correctly", func() {
			gomega.Expect(absoluteValueThenSignKey(-3, 1)).To(gomega.BeEquivalentTo(OrderingGreaterThan))
			gomega.Expect(absoluteValueThenSignKey(-3, -1)).To(gomega.BeEquivalentTo(OrderingGreaterThan))
			gomega.Expect(absoluteValueThenSignKey(3, 1)).To(gomega.BeEquivalentTo(OrderingGreaterThan))
			gomega.Expect(absoluteValueThenSignKey(3, -1)).To(gomega.BeEquivalentTo(OrderingGreaterThan))
			gomega.Expect(absoluteValueThenSignKey(-3, 3)).To(gomega.BeEquivalentTo(OrderingLessThan))
			gomega.Expect(absoluteValueThenSignKey(3, -3)).To(gomega.BeEquivalentTo(OrderingGreaterThan))
		})
		It("reverses comparison", func() {
			desc := CompareReverse(CompareOrdered[int])
			gomega.Expect(desc(-3, 1)).To(gomega.BeEquivalentTo(OrderingGreaterThan))
			gomega.Expect(desc(4, 4)).To(gomega.BeEquivalentTo(OrderingEqual))
			gomega.Expect(desc(3, 8)).To(gomega.BeEquivalentTo(OrderingGreaterThan))
			gomega.Expect(desc(8, 3)).To(gomega.BeEquivalentTo(OrderingLessThan))
		})
	})
}
