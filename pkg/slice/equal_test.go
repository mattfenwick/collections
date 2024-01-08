package slice

import (
	. "github.com/mattfenwick/collections/pkg/base"
	"github.com/mattfenwick/collections/pkg/function"
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func RunEqualTests() {
	Describe("Eq", func() {
		p1 := NewPair[Int, Bool](13, true)
		p2 := NewPair[Int, Bool](14, true)
		p3 := NewPair[Int, Bool](13, false)

		It("EqualBy", func() {
			gomega.Expect(EqualBy(function.On(Equal[Int], Fst[Int, Bool]))(p1, p2)).To(gomega.BeFalse())
			gomega.Expect(EqualBy(function.On(Equal[Int], Fst[Int, Bool]))(p1, p3)).To(gomega.BeTrue())
			gomega.Expect(EqualBy(function.On(Equal[Bool], Snd[Int, Bool]))(p1, p2)).To(gomega.BeTrue())
		})

		It("slice equality", func() {
			equal := EqualSlicePairwise[int]()
			gomega.Expect(equal([]int{}, []int{})).To(gomega.Equal(true))
			gomega.Expect(equal([]int{}, []int{18, 37})).To(gomega.Equal(false))
			gomega.Expect(equal([]int{25, 39}, []int{})).To(gomega.Equal(false))
			gomega.Expect(equal([]int{14, 32, 65, 8}, []int{14, 32, 65, 8})).To(gomega.Equal(true))
		})
		It("slice equality -- SliceOrd", func() {
			gomega.Expect(Equal[SliceOrd[Int]]([]Int{}, []Int{})).To(gomega.Equal(true))
		})
	})
}
