package dict

import (
	. "github.com/mattfenwick/collections/pkg/base"
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func RunEqualTests() {
	Describe("Eq", func() {
		p1 := map[string]int{"a": 1, "b": 2}
		p2 := map[string]int{"a": 2, "b": 2}
		p3 := map[string]int{"a": 2, "b": 1}
		//p4 := map[string]int{"a": 2, "b": 4}
		//p5 := map[string]int{"a": 1}
		//p6 := map[string]int{"b": 2}
		p7 := map[string]int{"a": 1, "b": 2, "c": 3}

		It("EqualMapPairwise", func() {
			equal := EqualMapPairwise[string, int]()
			gomega.Expect(equal(p1, p1)).To(gomega.Equal(true))
			gomega.Expect(equal(p1, p2)).To(gomega.Equal(false))
			gomega.Expect(equal(p1, p3)).To(gomega.Equal(false))

			gomega.Expect(equal(p2, p1)).To(gomega.Equal(false))
			gomega.Expect(equal(p2, p2)).To(gomega.Equal(true))
			gomega.Expect(equal(p2, p3)).To(gomega.Equal(false))

			gomega.Expect(equal(p3, p1)).To(gomega.Equal(false))
			gomega.Expect(equal(p3, p2)).To(gomega.Equal(false))
			gomega.Expect(equal(p3, p3)).To(gomega.Equal(true))

			gomega.Expect(equal(p7, p7)).To(gomega.Equal(true))
		})

		It("EqualMapIndex", func() {
			equalA := EqualMapIndex[string, int]("a")
			//equalB := maps.EqualMapIndex[string, int]("b")
			//equalC := maps.EqualMapIndex[string, int]("c")
			tests := []struct {
				E      Equaler[map[string]int]
				M1     map[string]int
				M2     map[string]int
				Result bool
			}{
				{E: equalA, M1: p1, M2: p1, Result: true},
			}
			for _, tc := range tests {
				gomega.Expect(tc.E(tc.M1, tc.M2)).To(gomega.Equal(tc.Result))
			}
			gomega.Expect(equalA(p1, p1)).To(gomega.BeTrue())
		})

	})
}
