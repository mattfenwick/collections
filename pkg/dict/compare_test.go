package dict

import (
	. "github.com/mattfenwick/collections/pkg/base"
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

//	func absoluteValue(i int) int {
//		if i < 0 {
//			return i * -1
//		}
//		return i
//	}
//
//	func isPositive[A builtin.Number](a A) bool {
//		return a > 0
//	}
//
// var absoluteValueThenSignKey = CompareBy(
//
//	function.On(builtin.CompareOrdered[int], absoluteValue),
//	function.On(builtin.CompareBool, isPositive[int]))
//
// var signThenAbsoluteValueKey = CompareBy(
//
//	function.On(builtin.CompareBool, isPositive[int]),
//	function.On(builtin.CompareOrdered[int], absoluteValue))
func RunCompareTests() {
	Describe("Compare", func() {
		p1 := map[string]int{"a": 1, "b": 2}
		p2 := map[string]int{"a": 2, "b": 2}
		p3 := map[string]int{"a": 2, "b": 1}
		p4 := map[string]int{"a": 2, "b": 4}
		p5 := map[string]int{"a": 1}
		p6 := map[string]int{"a": 2, "b": 2, "c": 1}
		p7 := map[string]int{"a": 1, "b": 2, "c": 3}
		p8 := map[string]int{"a": 1, "b": 3}

		It("map index ordering", func() {
			a := CompareMapIndex[string, int]("a")
			gomega.Expect(a(p1, p1)).To(gomega.BeEquivalentTo(OrderingEqual))
			gomega.Expect(a(p1, p2)).To(gomega.BeEquivalentTo(OrderingLessThan))
			gomega.Expect(a(p1, p3)).To(gomega.BeEquivalentTo(OrderingLessThan))
			gomega.Expect(a(p1, p5)).To(gomega.BeEquivalentTo(OrderingEqual))

			b := CompareMapIndex[string, int]("b")
			gomega.Expect(b(p2, p1)).To(gomega.BeEquivalentTo(OrderingEqual))
			gomega.Expect(b(p2, p3)).To(gomega.BeEquivalentTo(OrderingGreaterThan))
			gomega.Expect(b(p2, p4)).To(gomega.BeEquivalentTo(OrderingLessThan))

			c := CompareMapIndex[string, int]("c")
			gomega.Expect(c(p7, p1)).To(gomega.BeEquivalentTo(OrderingGreaterThan))
			gomega.Expect(c(p7, p4)).To(gomega.BeEquivalentTo(OrderingGreaterThan))
		})

		It("map pairwise ordering", func() {
			compare := CompareMapPairwise[string, int]()
			gomega.Expect(compare(p2, p2)).To(gomega.BeEquivalentTo(OrderingEqual))
			gomega.Expect(compare(p2, p1)).To(gomega.BeEquivalentTo(OrderingGreaterThan))
			gomega.Expect(compare(p2, p3)).To(gomega.BeEquivalentTo(OrderingGreaterThan))
			gomega.Expect(compare(p2, p4)).To(gomega.BeEquivalentTo(OrderingLessThan))
			gomega.Expect(compare(p2, p6)).To(gomega.BeEquivalentTo(OrderingLessThan))

			// extra key?  always greater, assuming prev kvs are all equal
			gomega.Expect(compare(p7, p1)).To(gomega.BeEquivalentTo(OrderingGreaterThan))
			gomega.Expect(compare(p7, p8)).To(gomega.BeEquivalentTo(OrderingLessThan))
			gomega.Expect(compare(p7, p3)).To(gomega.BeEquivalentTo(OrderingLessThan))
		})
	})

}
