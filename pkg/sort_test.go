package pkg

import (
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

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
}
