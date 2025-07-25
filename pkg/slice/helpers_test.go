package slice

import (
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func RunHelpersTests() {
	Describe("FilterMap", func() {
		in := []int{1, 2, 3}
		f := func(i int) *int {
			if i < 3 {
				return &i
			}
			return nil
		}
		It("Removes nils", func() {
			gomega.Expect(FilterMap(f, in)).To(gomega.Equal([]int{1, 2}))
		})
	})
}
