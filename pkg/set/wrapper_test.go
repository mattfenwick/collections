package set

import (
	"github.com/mattfenwick/collections/pkg/slice"
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func RunWrapperTests() {
	Describe("Wrapper", func() {
		s1 := NewWrapper([]int{})
		s2 := NewWrapper([]int{24, 3})
		s3 := NewWrapper([]int{8, 7, 8, 9})
		s4 := NewWrapper([]int{12, -12, 21, 3, 2})
		It("Len", func() {
			gomega.Expect(s1.Len()).To(gomega.Equal(0))
			gomega.Expect(s2.Len()).To(gomega.Equal(2))
			gomega.Expect(s3.Len()).To(gomega.Equal(3))
			gomega.Expect(s4.Len()).To(gomega.Equal(5))
		})

		It("basic methods", func() {
			ints := []int{13, 4, 12}
			w := NewWrapper(ints)
			gomega.Expect(w.Len()).To(gomega.Equal(3))
			for _, x := range ints {
				gomega.Expect(w.Contains(x)).To(gomega.BeTrue())
			}
			for _, x := range []int{-13, -4, 5, 3} {
				gomega.Expect(w.Contains(x)).To(gomega.BeFalse())
			}
			gomega.Expect(slice.Sort(w.ToSlice())).To(gomega.Equal([]int{4, 12, 13}))

			gomega.Expect(w.Delete(5)).To(gomega.BeFalse())
			gomega.Expect(w.Len()).To(gomega.Equal(3))

			gomega.Expect(w.Delete(12)).To(gomega.BeTrue())
			gomega.Expect(w.Len()).To(gomega.Equal(2))
			gomega.Expect(slice.Sort(w.ToSlice())).To(gomega.Equal([]int{4, 13}))

			gomega.Expect(w.Add(45)).To(gomega.BeTrue())
			gomega.Expect(w.Add(45)).To(gomega.BeFalse())
			gomega.Expect(w.Len()).To(gomega.Equal(3))
			gomega.Expect(slice.Sort(w.ToSlice())).To(gomega.Equal([]int{4, 13, 45}))
		})
	})
}
