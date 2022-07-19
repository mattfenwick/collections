package set

import (
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func RunSetTests() {
	Describe("Wrapper", func() {
		s1 := NewWrapper([]int{})
		s2 := NewWrapper([]int{24, 3})
		s3 := NewWrapper([]int{8, 7, 8, 9})
		s4 := NewWrapper([]int{12, -12, 21, 3, 2})
		It("", func() {
			gomega.Expect(s1.Len()).To(gomega.Equal(0))
			gomega.Expect(s2.Len()).To(gomega.Equal(2))
			gomega.Expect(s3.Len()).To(gomega.Equal(3))
			gomega.Expect(s4.Len()).To(gomega.Equal(5))
		})
	})

	Describe("Set", func() {
		s1 := NewSet([]int{})
		s2 := NewSet([]int{24, 3})
		s3 := NewSet([]int{8, 7, 8, 9})
		s4 := NewSet([]int{12, -12, 21, 3, 2})
		It("", func() {
			gomega.Expect(s1.Len()).To(gomega.Equal(0))
			gomega.Expect(s2.Len()).To(gomega.Equal(2))
			gomega.Expect(s3.Len()).To(gomega.Equal(3))
			gomega.Expect(s4.Len()).To(gomega.Equal(5))
		})
	})

}
