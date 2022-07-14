package pkg

import (
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func makePair[A Eq[A], B Eq[B]](a A, b B) PairEq[A, B] {
	return PairEq[A, B](*NewPair(a, b))
}

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
		It("Pair", func() {
			p1 := makePair[Int, Bool](13, true)
			p2 := makePair[Int, Bool](14, true)
			p3 := makePair[Int, Bool](13, false)
			gomega.Expect(Equal(&p1, &p1)).To(gomega.Equal(true))
			gomega.Expect(Equal(&p1, &p2)).To(gomega.Equal(false))
			gomega.Expect(Equal(&p1, &p3)).To(gomega.Equal(false))
			gomega.Expect(Equal(&p2, &p1)).To(gomega.Equal(false))
			gomega.Expect(Equal(&p2, &p2)).To(gomega.Equal(true))
			gomega.Expect(Equal(&p2, &p3)).To(gomega.Equal(false))
			gomega.Expect(Equal(&p3, &p1)).To(gomega.Equal(false))
			gomega.Expect(Equal(&p3, &p2)).To(gomega.Equal(false))
			gomega.Expect(Equal(&p3, &p3)).To(gomega.Equal(true))
		})
	})
}
