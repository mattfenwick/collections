package slices

import (
	. "github.com/mattfenwick/collections/pkg/base"
	"github.com/mattfenwick/collections/pkg/builtins"
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func RunCompareTests() {
	Describe("Compare", func() {
		It("slice ordering", func() {
			gomega.Expect(CompareSlice(builtins.Compare[int], []int{}, []int{})).To(gomega.BeEquivalentTo(OrderingEqual))

			gomega.Expect(CompareSlice(builtins.Compare[int], []int{4}, []int{})).To(gomega.BeEquivalentTo(OrderingGreaterThan))
			gomega.Expect(CompareSlice(builtins.Compare[int], []int{}, []int{4})).To(gomega.BeEquivalentTo(OrderingLessThan))
			gomega.Expect(CompareSlice(builtins.Compare[int], []int{4}, []int{4})).To(gomega.BeEquivalentTo(OrderingEqual))

			gomega.Expect(CompareSlice(builtins.Compare[int], []int{4, 7}, []int{4, 8})).To(gomega.BeEquivalentTo(OrderingLessThan))
			gomega.Expect(CompareSlice(builtins.Compare[int], []int{4, 7}, []int{4, 6})).To(gomega.BeEquivalentTo(OrderingGreaterThan))
			gomega.Expect(CompareSlice(builtins.Compare[int], []int{4, 7}, []int{4, 7})).To(gomega.BeEquivalentTo(OrderingEqual))

			gomega.Expect(CompareSlice(builtins.Compare[int], []int{1, 2, 3}, []int{3, 4, 5})).To(gomega.BeEquivalentTo(OrderingLessThan))
			gomega.Expect(CompareSlice(builtins.Compare[int], []int{1, 2, 3}, []int{3})).To(gomega.BeEquivalentTo(OrderingLessThan))
			gomega.Expect(CompareSlice(builtins.Compare[int], []int{1, 2, 3}, []int{})).To(gomega.BeEquivalentTo(OrderingGreaterThan))

			gomega.Expect(CompareSlice(builtins.Compare[int], []int{3, 4, 5}, []int{1, 2, 3})).To(gomega.BeEquivalentTo(OrderingGreaterThan))
			gomega.Expect(CompareSlice(builtins.Compare[int], []int{3}, []int{1, 2, 3})).To(gomega.BeEquivalentTo(OrderingGreaterThan))
			gomega.Expect(CompareSlice(builtins.Compare[int], []int{}, []int{1, 2, 3})).To(gomega.BeEquivalentTo(OrderingLessThan))
		})
	})

	Describe("Eq", func() {
		It("slice equality", func() {
			gomega.Expect(EqualSlice(builtins.Equal[int], []int{}, []int{})).To(gomega.Equal(true))
			gomega.Expect(EqualSlice(builtins.Equal[int], []int{}, []int{18, 37})).To(gomega.Equal(false))
			gomega.Expect(EqualSlice(builtins.Equal[int], []int{25, 39}, []int{})).To(gomega.Equal(false))
			gomega.Expect(EqualSlice(builtins.Equal[int], []int{14, 32, 65, 8}, []int{14, 32, 65, 8})).To(gomega.Equal(true))
		})
		It("slice equality -- SliceOrd", func() {
			gomega.Expect(Equal[SliceOrd[Int]]([]Int{}, []Int{})).To(gomega.Equal(true))
		})
	})
}
