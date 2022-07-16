package slices

import (
	"github.com/mattfenwick/collections/pkg"
	"github.com/mattfenwick/collections/pkg/base"
	"github.com/mattfenwick/collections/pkg/builtins"
	"github.com/mattfenwick/collections/pkg/functions"
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func RunPreludeTests() {
	Describe("Special folds", func() {
		It("Maximum", func() {
			gomega.Expect(Maximum[base.Int]([]base.Int{})).To(gomega.Equal(pkg.Nothing[base.Int]()))
			gomega.Expect(Maximum[base.Int]([]base.Int{3})).To(gomega.Equal(pkg.Just[base.Int](3)))
			gomega.Expect(Maximum[base.Int]([]base.Int{3, 2, -3})).To(gomega.Equal(pkg.Just[base.Int](3)))
			gomega.Expect(Maximum[base.Int]([]base.Int{-2, 3, 2})).To(gomega.Equal(pkg.Just[base.Int](3)))
		})
	})

	Describe("Scans", func() {
		It("Scanl", func() {
			gomega.Expect(
				Scanl(
					builtins.Plus[int],
					0,
					Range(1, 5, 1))).
				To(gomega.Equal([]int{0, 1, 3, 6, 10}))

			gomega.Expect(
				Scanl(
					builtins.Plus[int],
					42,
					[]int{})).
				To(gomega.Equal([]int{42}))

			gomega.Expect(
				Scanl(
					builtins.Minus[int],
					100,
					Range(1, 5, 1))).
				To(gomega.Equal([]int{100, 99, 97, 94, 90}))

			gomega.Expect(
				Scanl(
					func(state []int, next int) []int { return Cons(next, state) },
					[]int{5, 3, 3},
					Range(6, 10, 1))).
				To(gomega.Equal([][]int{
					{5, 3, 3},
					{6, 5, 3, 3},
					{7, 6, 5, 3, 3},
					{8, 7, 6, 5, 3, 3},
					{9, 8, 7, 6, 5, 3, 3},
				}))
		})
		It("Scanl1", func() {
			gomega.Expect(
				Scanl1(
					builtins.Plus[int],
					Range(1, 5, 1))).
				To(gomega.Equal([]int{1, 3, 6, 10}))

			gomega.Expect(
				Scanl1(
					builtins.Plus[int],
					[]int{})).
				To(gomega.Equal([]int{}))

			gomega.Expect(
				Scanl1(
					builtins.Plus[int],
					[]int{42})).
				To(gomega.Equal([]int{42}))

			gomega.Expect(
				Scanl1(
					builtins.Minus[int],
					Range(1, 5, 1))).
				To(gomega.Equal([]int{1, -1, -4, -8}))

			gomega.Expect(
				Scanl1(
					builtins.And,
					[]bool{true, false, true, false})).
				To(gomega.Equal([]bool{true, false, false, false}))

			gomega.Expect(
				Scanl1(
					builtins.Or,
					[]bool{false, true, false, true})).
				To(gomega.Equal([]bool{false, true, true, true}))
		})
		It("Scanr", func() {
			gomega.Expect(
				Scanr(
					builtins.Plus[int],
					0,
					Range(1, 5, 1))).
				To(gomega.Equal([]int{10, 9, 7, 4, 0}))

			gomega.Expect(
				Scanr(
					builtins.Plus[int],
					42,
					[]int{})).
				To(gomega.Equal([]int{42}))

			gomega.Expect(
				Scanr(
					builtins.Minus[int],
					100,
					Range(1, 5, 1))).
				To(gomega.Equal([]int{98, -97, 99, -96, 100}))

			gomega.Expect(
				Scanr(
					Cons[int],
					[]int{5, 3, 3},
					Range(6, 10, 1))).
				To(gomega.Equal([][]int{
					{6, 7, 8, 9, 5, 3, 3},
					{7, 8, 9, 5, 3, 3},
					{8, 9, 5, 3, 3},
					{9, 5, 3, 3},
					{5, 3, 3},
				}))
		})
		It("Scanr1", func() {
			gomega.Expect(
				Scanr1(
					builtins.Plus[int],
					Range(1, 5, 1))).
				To(gomega.Equal([]int{10, 9, 7, 4}))

			gomega.Expect(
				Scanr1(
					builtins.Plus[int],
					[]int{})).
				To(gomega.Equal([]int{}))

			gomega.Expect(
				Scanr1(
					builtins.Plus[int],
					[]int{42})).
				To(gomega.Equal([]int{42}))

			gomega.Expect(
				Scanr1(
					builtins.Minus[int],
					Range(1, 5, 1))).
				To(gomega.Equal([]int{-2, 3, -1, 4}))

			gomega.Expect(
				Scanr1(
					builtins.And,
					[]bool{true, false, true, true})).
				To(gomega.Equal([]bool{false, false, true, true}))

			gomega.Expect(
				Scanr1(
					builtins.Or,
					[]bool{false, true, false, false})).
				To(gomega.Equal([]bool{true, true, false, false}))
		})
	})

	Describe("Accumulating maps", func() {
		It("MapAccumL", func() {
			gomega.Expect(MapAccumL(
				func(a int, b int) *base.Pair[int, int] {
					return base.NewPair[int, int](a+b, a)
				},
				0,
				Range[int](1, 11, 1),
			)).To(gomega.Equal(base.NewPair[int, []int](55, []int{0, 1, 3, 6, 10, 15, 21, 28, 36, 45})))

			gomega.Expect(MapAccumL(
				func(b []int, a int) *base.Pair[[]int, []int] {
					return base.NewPair[[]int, []int](Append(b, []int{a}), b)
				},
				[]int{0},
				Range[int](1, 5, 1),
			)).To(gomega.Equal(base.NewPair[[]int, [][]int](Range(0, 5, 1), [][]int{
				Range(0, 1, 1),
				Range(0, 2, 1),
				Range(0, 3, 1),
				Range(0, 4, 1),
			})))
		})
		It("MapAccumR", func() {
			gomega.Expect(MapAccumR(
				func(a int, b int) *base.Pair[int, int] {
					return base.NewPair[int, int](a+b, a)
				},
				0,
				Range[int](1, 11, 1),
			)).To(gomega.Equal(base.NewPair[int, []int](55, []int{54, 52, 49, 45, 40, 34, 27, 19, 10, 0})))

			gomega.Expect(MapAccumR(
				func(b []int, a int) *base.Pair[[]int, []int] {
					return base.NewPair[[]int, []int](Append(b, []int{a}), b)
				},
				[]int{0},
				Range[int](1, 5, 1),
			)).To(gomega.Equal(base.NewPair[[]int, [][]int]([]int{0, 4, 3, 2, 1}, [][]int{
				{0, 4, 3, 2},
				{0, 4, 3},
				{0, 4},
				{0},
			})))
		})
	})

	Describe("Infinite lists", func() {
		It("Iterate", func() {
			gomega.Expect(Iterate(5, builtins.Not, true)).To(gomega.Equal([]bool{true, false, true, false, true}))
			gomega.Expect(Iterate(5, functions.Partial2(builtins.Plus[int])(3), 42)).To(gomega.Equal([]int{42, 45, 48, 51, 54}))
		})
	})

	Describe("Unfolding", func() {
		It("Unfoldr", func() {
			gomega.Expect(
				Unfoldr(func(next int) *pkg.Maybe[*base.Pair[int, int]] {
					if next == 0 {
						return pkg.Nothing[*base.Pair[int, int]]()
					}
					return pkg.Just(base.NewPair(next, next-1))
				}, 10)).
				To(gomega.Equal([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}))
		})
	})

	eq3 := functions.Partial2(builtins.EQ[int])(3)

	Describe("Sublists", func() {
		lt3 := functions.Partial2(functions.Flip(builtins.LT[int]))(3)
		pair := func(xs []int, ys []int) *base.Pair[[]int, []int] {
			return base.NewPair(xs, ys)
		}
		It("TakeWhile", func() {
			gomega.Expect(TakeWhile(lt3, []int{1, 2, 3, 4, 1, 2, 3, 4})).To(gomega.Equal([]int{1, 2}))
			gomega.Expect(TakeWhile(lt3, []int{-18, 2, 0, 1})).To(gomega.Equal([]int{-18, 2, 0, 1}))
			gomega.Expect(TakeWhile(lt3, []int{8, 1, 4})).To(gomega.BeNil())
			gomega.Expect(TakeWhile(eq3, []int{3, 2, 4, 3, 1})).To(gomega.Equal([]int{3}))
		})
		It("DropWhile", func() {
			gomega.Expect(DropWhile(lt3, []int{1, 2, 3, 4, 1, 2, 3, 4})).To(gomega.Equal([]int{3, 4, 1, 2, 3, 4}))
			gomega.Expect(DropWhile(lt3, []int{-18, 2, 0, 1})).To(gomega.BeNil())
			gomega.Expect(DropWhile(lt3, []int{8, 1, 4})).To(gomega.Equal([]int{8, 1, 4}))
			gomega.Expect(DropWhile(eq3, []int{3, 2, 4, 3, 1})).To(gomega.Equal([]int{2, 4, 3, 1}))
		})
		It("Span", func() {
			gomega.Expect(Span(lt3, []int{1, 2, 3, 4, 1, 2, 3, 4})).To(gomega.Equal(pair([]int{1, 2}, []int{3, 4, 1, 2, 3, 4})))
			gomega.Expect(Span(lt3, []int{-18, 2, 0, 1})).To(gomega.Equal(pair([]int{-18, 2, 0, 1}, nil)))
			gomega.Expect(Span(lt3, []int{8, 1, 4})).To(gomega.Equal(pair(nil, []int{8, 1, 4})))
			gomega.Expect(Span[int](eq3, []int{3, 2, 4, 3, 1})).To(gomega.Equal(pair([]int{3}, []int{2, 4, 3, 1})))
		})
	})

	Describe("the by operations", func() {
		It("GroupBy", func() {
			gomega.Expect(GroupBy(builtins.EQ[int], []int{})).To(gomega.BeNil())
			gomega.Expect(GroupBy(builtins.EQ[int], []int{1})).To(gomega.Equal([][]int{{1}}))
			gomega.Expect(GroupBy(builtins.EQ[int], []int{1, 1, 1})).To(gomega.Equal([][]int{{1, 1, 1}}))
			gomega.Expect(GroupBy(builtins.EQ[int], []int{1, 2, 1})).To(gomega.Equal([][]int{{1}, {2}, {1}}))
			gomega.Expect(GroupBy(builtins.EQ[int], []int{2, 1, 1})).To(gomega.Equal([][]int{{2}, {1, 1}}))
			gomega.Expect(GroupBy(builtins.EQ[int], []int{7, 4, 2, 2, 4, 2, 2, 4, 9, 9, 4})).To(gomega.Equal([][]int{
				{7},
				{4},
				{2, 2},
				{4},
				{2, 2},
				{4},
				{9, 9},
				{4},
			}))
		})
	})
}
