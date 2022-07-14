package pkg

import (
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func absoluteValue(i Int) Int {
	if i < 0 {
		return i * -1
	}
	return i
}

func intKey(i Int) PairOrd[Int, Bool] {
	var isPositive Bool = false
	if i > 0 {
		isPositive = true
	}
	return PairOrd[Int, Bool](*NewPair[Int, Bool](absoluteValue(i), isPositive))
}

func RunSortTests() {
	Describe("Sort", func() {
		It("Empty", func() {
			gomega.Expect(Sort([]String{})).To(gomega.Equal([]String{}))
		})
		It("Single element", func() {
			gomega.Expect(Sort([]String{"3"})).To(gomega.Equal([]String{"3"}))
		})
		It("Already sorted", func() {
			gomega.Expect(Sort([]Int{-3, 12, 37, 45, 59})).To(gomega.Equal([]Int{-3, 12, 37, 45, 59}))
		})
		It("Duplicate elements", func() {
			gomega.Expect(Sort([]Int{2, 1, 2, 5, 2, 0, 2})).To(gomega.Equal([]Int{0, 1, 2, 2, 2, 2, 5}))
		})
	})
	Describe("SortOn", func() {
		ints := []Int{1, 18, -34, 79, 97, 36, 42, -18, -3, -1, -18}
		It("key of self", func() {
			sorted := SortOn(ints, Id[Int])
			gomega.Expect(sorted).To(gomega.Equal([]Int{-34, -18, -18, -3, -1, 1, 18, 36, 42, 79, 97}))
		})
		It("key of Pair", func() {
			sorted := SortOn(ints, intKey)
			gomega.Expect(sorted).To(gomega.Equal([]Int{-1, 1, -3, -18, -18, 18, -34, 36, 42, 79, 97}))
		})
	})
}
