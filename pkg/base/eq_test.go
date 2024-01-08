package base

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

		It("EqualPairEq", func() {
			p1 := NewPair[Int, Bool](13, true)
			p2 := NewPair[Int, Bool](14, true)
			p3 := NewPair[Int, Bool](13, false)
			equal := EqualPairEq[Int, Bool]()

			gomega.Expect(equal(p1, p1)).To(gomega.Equal(true))
			gomega.Expect(equal(p1, p2)).To(gomega.Equal(false))
			gomega.Expect(equal(p1, p3)).To(gomega.Equal(false))
			gomega.Expect(equal(p2, p1)).To(gomega.Equal(false))
			gomega.Expect(equal(p2, p2)).To(gomega.Equal(true))
			gomega.Expect(equal(p2, p3)).To(gomega.Equal(false))
			gomega.Expect(equal(p3, p1)).To(gomega.Equal(false))
			gomega.Expect(equal(p3, p2)).To(gomega.Equal(false))
			gomega.Expect(equal(p3, p3)).To(gomega.Equal(true))
		})

		It("EqualPair", func() {
			p1 := NewPair[int, bool](13, true)
			p2 := NewPair[int, bool](14, true)
			p3 := NewPair[int, bool](13, false)
			equal := EqualPair[int, bool]()

			gomega.Expect(equal(p1, p1)).To(gomega.Equal(true))
			gomega.Expect(equal(p1, p2)).To(gomega.Equal(false))
			gomega.Expect(equal(p1, p3)).To(gomega.Equal(false))
			gomega.Expect(equal(p2, p1)).To(gomega.Equal(false))
			gomega.Expect(equal(p2, p2)).To(gomega.Equal(true))
			gomega.Expect(equal(p2, p3)).To(gomega.Equal(false))
			gomega.Expect(equal(p3, p1)).To(gomega.Equal(false))
			gomega.Expect(equal(p3, p2)).To(gomega.Equal(false))
			gomega.Expect(equal(p3, p3)).To(gomega.Equal(true))
		})
	})
}
