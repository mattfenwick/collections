package set

import (
	"github.com/mattfenwick/collections/pkg/iterable"
	"github.com/mattfenwick/collections/pkg/slice"
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func RunSetTests() {
	Describe("Set", func() {
		s1 := FromSlice([]int{})
		s2 := FromSlice([]int{24, 3})
		s3 := FromSlice([]int{8, 7, 8, 9})
		s4 := FromSlice([]int{12, -12, 21, 3, 2})

		It("basic methods", func() {
			ints := []int{13, 4, 12}
			s := FromSlice(ints)
			gomega.Expect(s.Len()).To(gomega.Equal(3))
			for _, x := range ints {
				gomega.Expect(s.Contains(x)).To(gomega.BeTrue())
			}
			for _, x := range []int{-13, -4, 5, 3} {
				gomega.Expect(s.Contains(x)).To(gomega.BeFalse())
			}
			gomega.Expect(slice.Sort(s.ToSlice())).To(gomega.Equal([]int{4, 12, 13}))

			gomega.Expect(s.Delete(5)).To(gomega.BeFalse())
			gomega.Expect(s.Len()).To(gomega.Equal(3))

			gomega.Expect(s.Delete(12)).To(gomega.BeTrue())
			gomega.Expect(s.Len()).To(gomega.Equal(2))
			gomega.Expect(slice.Sort(s.ToSlice())).To(gomega.Equal([]int{4, 13}))

			gomega.Expect(s.Add(45)).To(gomega.BeTrue())
			gomega.Expect(s.Add(45)).To(gomega.BeFalse())
			gomega.Expect(s.Len()).To(gomega.Equal(3))
			gomega.Expect(slice.Sort(s.ToSlice())).To(gomega.Equal([]int{4, 13, 45}))
		})

		It("Len", func() {
			gomega.Expect(s1.Len()).To(gomega.Equal(0))
			gomega.Expect(s2.Len()).To(gomega.Equal(2))
			gomega.Expect(s3.Len()).To(gomega.Equal(3))
			gomega.Expect(s4.Len()).To(gomega.Equal(5))
		})

		It("Union", func() {
			set := FromSlice([]int{4, 3, 2, 1})
			union := set.Union(s4)
			gomega.Expect(union.Len()).To(gomega.Equal(7))
			gomega.Expect(slice.Sort(union.ToSlice())).To(gomega.Equal([]int{-12, 1, 2, 3, 4, 12, 21}))
		})

		It("Intersection", func() {
			set := FromSlice([]int{4, 3, 2, 1})
			intersection := set.Intersect(s4)
			gomega.Expect(intersection.Len()).To(gomega.Equal(2))
			gomega.Expect(slice.Sort(intersection.ToSlice())).To(gomega.Equal([]int{2, 3}))
		})

		It("Difference", func() {
			set := FromSlice([]int{4, 3, 2, 1})
			diff := set.Difference(s4)
			gomega.Expect(diff.Len()).To(gomega.Equal(2))
			gomega.Expect(slice.Sort(diff.ToSlice())).To(gomega.Equal([]int{1, 4}))
		})

		It("Iterator", func() {
			ints := slice.Sort(iterable.ToSlice(s4.Iterator()))
			gomega.Expect(ints).To(gomega.Equal([]int{-12, 2, 3, 12, 21}))
		})
	})
}
