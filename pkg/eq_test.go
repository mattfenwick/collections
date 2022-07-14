package pkg

import (
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func RunEqTests() {
	Describe("Eq", func() {
		It("Int", func() {
			gomega.Expect(Equal[Int](13, 35)).To(gomega.Equal(false))
			gomega.Expect(Equal[Int](13, 3)).To(gomega.Equal(false))
			gomega.Expect(Equal[Int](13, 13)).To(gomega.Equal(true))
		})
		It("Bool", func() {
			gomega.Expect(Equal[Bool](true, false)).To(gomega.Equal(false))
			gomega.Expect(Equal[Bool](false, true)).To(gomega.Equal(false))
			gomega.Expect(Equal[Bool](false, false)).To(gomega.Equal(true))
			gomega.Expect(Equal[Bool](true, true)).To(gomega.Equal(true))
		})
		It("SliceEq", func() {
			gomega.Expect(Equal[SliceEq[Int]]([]Int{}, []Int{})).To(gomega.Equal(true))
			gomega.Expect(Equal[SliceEq[Int]]([]Int{}, []Int{18, 37})).To(gomega.Equal(false))
			gomega.Expect(Equal[SliceEq[Int]]([]Int{25, 39}, []Int{})).To(gomega.Equal(false))
			gomega.Expect(Equal[SliceEq[Int]]([]Int{14, 32, 65, 8}, []Int{14, 32, 65, 8})).To(gomega.Equal(true))
		})
		It("SliceOrd", func() {
			gomega.Expect(Equal[SliceOrd[Int]]([]Int{}, []Int{})).To(gomega.Equal(true))
		})
	})
}
