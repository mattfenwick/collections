package slices

import (
	"github.com/mattfenwick/collections/pkg"
	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

func RunSliceTests() {
	Describe("Scan", func() {
		It("Scanl", func() {
			gomega.Expect(
				Scanl(
					pkg.Plus[int],
					0,
					Range(1, 5, 1))).
				To(gomega.Equal([]int{0, 1, 3, 6, 10}))

			gomega.Expect(
				Scanl(
					pkg.Plus[int],
					42,
					[]int{})).
				To(gomega.Equal([]int{42}))

			gomega.Expect(
				Scanl(
					pkg.Minus[int],
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
					pkg.Plus[int],
					Range(1, 5, 1))).
				To(gomega.Equal([]int{1, 3, 6, 10}))

			gomega.Expect(
				Scanl1(
					pkg.Plus[int],
					[]int{})).
				To(gomega.Equal([]int{}))

			gomega.Expect(
				Scanl1(
					pkg.Plus[int],
					[]int{42})).
				To(gomega.Equal([]int{42}))

			gomega.Expect(
				Scanl1(
					pkg.Minus[int],
					Range(1, 5, 1))).
				To(gomega.Equal([]int{1, -1, -4, -8}))

			gomega.Expect(
				Scanl1(
					pkg.And,
					[]bool{true, false, true, false})).
				To(gomega.Equal([]bool{true, false, false, false}))

			gomega.Expect(
				Scanl1(
					pkg.Or,
					[]bool{false, true, false, true})).
				To(gomega.Equal([]bool{false, true, true, true}))
		})
		It("Scanr", func() {
			gomega.Expect(
				Scanr(
					pkg.Plus[int],
					0,
					Range(1, 5, 1))).
				To(gomega.Equal([]int{10, 9, 7, 4, 0}))

			gomega.Expect(
				Scanr(
					pkg.Plus[int],
					42,
					[]int{})).
				To(gomega.Equal([]int{42}))

			gomega.Expect(
				Scanr(
					pkg.Minus[int],
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
					pkg.Plus[int],
					Range(1, 5, 1))).
				To(gomega.Equal([]int{10, 9, 7, 4}))

			gomega.Expect(
				Scanr1(
					pkg.Plus[int],
					[]int{})).
				To(gomega.Equal([]int{}))

			gomega.Expect(
				Scanr1(
					pkg.Plus[int],
					[]int{42})).
				To(gomega.Equal([]int{42}))

			gomega.Expect(
				Scanr1(
					pkg.Minus[int],
					Range(1, 5, 1))).
				To(gomega.Equal([]int{-2, 3, -1, 4}))

			gomega.Expect(
				Scanr1(
					pkg.And,
					[]bool{true, false, true, true})).
				To(gomega.Equal([]bool{false, false, true, true}))

			gomega.Expect(
				Scanr1(
					pkg.Or,
					[]bool{false, true, false, false})).
				To(gomega.Equal([]bool{true, true, false, false}))
		})
	})
}
