package set

import (
	"github.com/mattfenwick/collections/pkg/slice"
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func RunSetTests() {
	Describe("Set", func() {
		s1 := NewSet([]int{})
		s2 := NewSet([]int{24, 3})
		s3 := NewSet([]int{8, 7, 8, 9})
		s4 := NewSet([]int{12, -12, 21, 3, 2})

		It("basic methods", func() {
			ints := []int{13, 4, 12}
			s := NewSet(ints)
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
			set := NewSet([]int{4, 3, 2, 1})
			gomega.Expect(set.Union(s4)).To(gomega.Equal(3))
			gomega.Expect(set.Len()).To(gomega.Equal(7))
			gomega.Expect(slice.Sort(set.ToSlice())).To(gomega.Equal([]int{-12, 1, 2, 3, 4, 12, 21}))
		})

		It("Intersection", func() {
			set := NewSet([]int{4, 3, 2, 1})
			gomega.Expect(set.Intersect(s4)).To(gomega.Equal(2))
			gomega.Expect(set.Len()).To(gomega.Equal(2))
			gomega.Expect(slice.Sort(set.ToSlice())).To(gomega.Equal([]int{2, 3}))
		})

		It("Difference", func() {
			set := NewSet([]int{4, 3, 2, 1})
			gomega.Expect(set.Difference(s4)).To(gomega.Equal(2))
			gomega.Expect(set.Len()).To(gomega.Equal(2))
			gomega.Expect(slice.Sort(set.ToSlice())).To(gomega.Equal([]int{1, 4}))
		})
	})
}
