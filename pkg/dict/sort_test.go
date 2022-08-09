package dict

import (
	"github.com/mattfenwick/collections/pkg/base"
	"github.com/mattfenwick/collections/pkg/slice"
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func RunSortTests() {
	ds := []map[string]int{
		{"d": 22, "m": 39},
		{},
		{"c": 10},
		{"c": 7},
		{"c": 13},
		{"a": 16},
		{"a": 2},
	}
	Describe("Sort", func() {
		It("CompareMapPairwise", func() {
			sorted := slice.SortBy(CompareMapPairwise[string, int](), ds)
			gomega.Expect(sorted).To(gomega.Equal([]map[string]int{
				{},
				{"a": 2},
				{"a": 16},
				{"c": 7},
				{"c": 10},
				{"c": 13},
				{"d": 22, "m": 39},
			}))
		})
		It("CompareMapIndex", func() {
			d1 := map[string]int{"a": 4}
			d2 := map[string]int{"a": 2, "b": 12}
			d3 := map[string]int{"b": 17}

			a := CompareMapIndex[string, int]("a")
			b := CompareMapIndex[string, int]("b")
			c := CompareMapIndex[string, int]("c")

			gomega.Expect(a(d1, d1)).To(gomega.BeEquivalentTo(base.OrderingEqual))
			gomega.Expect(a(d1, d2)).To(gomega.BeEquivalentTo(base.OrderingGreaterThan))
			gomega.Expect(a(d1, d3)).To(gomega.BeEquivalentTo(base.OrderingGreaterThan))

			gomega.Expect(b(d2, d1)).To(gomega.BeEquivalentTo(base.OrderingGreaterThan))
			gomega.Expect(b(d2, d2)).To(gomega.BeEquivalentTo(base.OrderingEqual))
			gomega.Expect(b(d2, d3)).To(gomega.BeEquivalentTo(base.OrderingLessThan))

			gomega.Expect(c(d3, d1)).To(gomega.BeEquivalentTo(base.OrderingEqual))
			gomega.Expect(c(d3, d2)).To(gomega.BeEquivalentTo(base.OrderingEqual))
			gomega.Expect(c(d3, d3)).To(gomega.BeEquivalentTo(base.OrderingEqual))
		})
	})
}
