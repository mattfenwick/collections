package dict

import (
	"github.com/mattfenwick/collections/pkg/slice"
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func RunDictTests() {
	Describe("basic dict functionality", func() {
		d := map[int]int{
			5:  7,
			11: 3,
			2:  4,
		}
		It("Keys", func() {
			gomega.Expect(slice.Sort(Keys(d))).To(gomega.Equal([]int{2, 5, 11}))
		})
		It("Values", func() {
			gomega.Expect(slice.Sort(Values(d))).To(gomega.Equal([]int{3, 4, 7}))
		})
	})
}
