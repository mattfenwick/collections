package dict

import (
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
	})
}
