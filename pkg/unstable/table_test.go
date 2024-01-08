package unstable

import (
	"github.com/mattfenwick/collections/pkg/base"
	"github.com/mattfenwick/collections/pkg/iterable"
	"github.com/mattfenwick/collections/pkg/slice"
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func pair[A any, B any](a A, b B) *base.Pair[A, B] {
	return base.NewPair(a, b)
}

func RunTableTests() {
	Describe("Table", func() {
		t1 := FromSlice([]*base.Pair[int, int]{})
		t2 := FromSlice([]*base.Pair[int, int]{pair(24, 12), pair(3, 5)})
		t3 := FromSlice([]*base.Pair[int, int]{pair(8, 1), pair(7, 2), pair(8, 3), pair(9, 4)})
		t4 := FromSlice([]*base.Pair[int, int]{pair(12, 1), pair(-12, 2), pair(21, 3), pair(3, 4), pair(2, 5)})
		empty := NewTable[int, int](nil)

		It("handles nils", func() {
			gomega.Expect(empty.Len()).To(gomega.Equal(0))
			gomega.Expect(empty.ToSlice()).To(gomega.Equal([]*base.Pair[int, int]{}))
		})

		It("basic methods", func() {
			ints := []*base.Pair[int, int]{pair(13, 1), pair(4, 2), pair(12, 3)}
			t := FromSlice(ints)
			gomega.Expect(t.Len()).To(gomega.Equal(3))
			for _, x := range ints {
				k, v := x.Fst, x.Snd
				gomega.Expect(t.Contains(k)).To(gomega.BeTrue())
				gomega.Expect(*t.Get(k)).To(gomega.Equal(v))
			}
			for _, x := range []int{-13, -4, 5, 3} {
				gomega.Expect(t.Contains(x)).To(gomega.BeFalse())
				gomega.Expect(t.Get(x)).To(gomega.BeNil())
			}
			gomega.Expect(slice.SortOn(base.Fst[int, int], t.ToSlice())).To(
				gomega.Equal([]*base.Pair[int, int]{pair(4, 2), pair(12, 3), pair(13, 1)}))

			gomega.Expect(t.Delete(5)).To(gomega.BeNil())
			gomega.Expect(t.Len()).To(gomega.Equal(3))

			gomega.Expect(*t.Delete(12)).To(gomega.Equal(3))
			gomega.Expect(t.Len()).To(gomega.Equal(2))
			gomega.Expect(slice.SortOn(base.Fst[int, int], t.ToSlice())).To(
				gomega.Equal([]*base.Pair[int, int]{pair(4, 2), pair(13, 1)}))

			gomega.Expect(t.Set(45, 4)).To(gomega.BeTrue())
			gomega.Expect(t.Set(45, 4)).To(gomega.BeFalse())
			gomega.Expect(t.Len()).To(gomega.Equal(3))
			gomega.Expect(slice.SortOn(base.Fst[int, int], t.ToSlice())).To(
				gomega.Equal([]*base.Pair[int, int]{pair(4, 2), pair(13, 1), pair(45, 4)}))
		})

		It("Len", func() {
			gomega.Expect(t1.Len()).To(gomega.Equal(0))
			gomega.Expect(t2.Len()).To(gomega.Equal(2))
			gomega.Expect(t3.Len()).To(gomega.Equal(3))
			gomega.Expect(t4.Len()).To(gomega.Equal(5))
		})

		It("Merge", func() {
			cases := []*struct {
				Left     *Table[int, int]
				Right    *Table[int, int]
				Expected []*base.Pair[int, int]
			}{
				{t2, empty, Entries(t2)},
				{empty, t2, Entries(t2)},
				{t2, t2, Entries(t2)},
				{t2, t4, []*base.Pair[int, int]{
					pair(-12, 2),
					pair(2, 5),
					pair(3, 4),
					pair(12, 1),
					pair(21, 3),
					pair(24, 12)}},
				{t4, t2, []*base.Pair[int, int]{
					pair(-12, 2),
					pair(2, 5),
					pair(3, 5),
					pair(12, 1),
					pair(21, 3),
					pair(24, 12)}},
			}
			for _, c := range cases {
				actual := Entries(c.Left.Merge(c.Right))
				gomega.Expect(actual).To(gomega.Equal(c.Expected))
			}
		})

		//It("Union", func() {
		//	set := FromSlice([]int{4, 3, 2, 1})
		//	union := set.Union(s4)
		//	gomega.Expect(union.Len()).To(gomega.Equal(7))
		//	gomega.Expect(slice.Sort(union.ToSlice())).To(gomega.Equal([]int{-12, 1, 2, 3, 4, 12, 21}))
		//})
		//
		//It("Intersection", func() {
		//	set := FromSlice([]int{4, 3, 2, 1})
		//	intersection := set.Intersect(s4)
		//	gomega.Expect(intersection.Len()).To(gomega.Equal(2))
		//	gomega.Expect(slice.Sort(intersection.ToSlice())).To(gomega.Equal([]int{2, 3}))
		//})
		//
		//It("Difference", func() {
		//	set := FromSlice([]int{4, 3, 2, 1})
		//	diff := set.Difference(s4)
		//	gomega.Expect(diff.Len()).To(gomega.Equal(2))
		//	gomega.Expect(slice.Sort(diff.ToSlice())).To(gomega.Equal([]int{1, 4}))
		//})

		It("Iterator", func() {
			ints := slice.SortOn(base.Fst[int, int], iterable.ToSlice(t4.Iterator()))
			gomega.Expect(ints).To(
				gomega.Equal([]*base.Pair[int, int]{pair(-12, 2), pair(2, 5), pair(3, 4), pair(12, 1), pair(21, 3)}))
		})

		It("When using a pointer type as a key: difference between using golang's builtin comparable vs. user-defined equality", func() {
			p1 := base.NewPair(base.NewPair(4, 5), 3)
			p2 := base.NewPair(base.NewPair(4, 5), 3)
			pairs := []*base.Pair[*base.Pair[int, int], int]{p1, p2}

			byBuiltin := FromSlice(pairs)
			gomega.Expect(byBuiltin.Len()).To(gomega.Equal(2))

			byUserDefined := FromSliceBy(base.Fst[int, int], pairs)
			gomega.Expect(byUserDefined.Len()).To(gomega.Equal(1))
		})
	})
}
