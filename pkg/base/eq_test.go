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
	})
}
