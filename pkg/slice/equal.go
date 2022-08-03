package slice

import (
	. "github.com/mattfenwick/collections/pkg/base"
	"github.com/mattfenwick/collections/pkg/builtin"
	"github.com/mattfenwick/collections/pkg/function"
	"golang.org/x/exp/slices"
)

func EqualSliceIndexEq[A Eq[A]](i int) Equaler[[]A] {
	return EqualSliceIndexBy(i, Equal[A])
}

func EqualSliceIndex[A comparable](i int) Equaler[[]A] {
	return EqualSliceIndexBy(i, builtin.Equal[A])
}

// EqualSliceIndexBy looks at a single index
func EqualSliceIndexBy[A any](i int, equal Equaler[A]) Equaler[[]A] {
	return func(xs []A, ys []A) bool {
		if i < len(xs) && i < len(ys) {
			return equal(xs[i], ys[i])
		} else if i >= len(xs) && i >= len(ys) {
			// ran off the end of both slices
			return true
		}
		// ran off the end of just one slice
		return false
	}
}

func EqualSlicePairwiseEq[A Eq[A]]() Equaler[[]A] {
	return EqualSlicePairwiseBy(Equal[A])
}

func EqualSlicePairwise[A comparable]() Equaler[[]A] {
	return EqualSlicePairwiseBy(builtin.Equal[A])
}

func EqualSlicePairwiseBy[A any](equal Equaler[A]) Equaler[[]A] {
	return func(xs []A, ys []A) bool {
		return slices.EqualFunc(xs, ys, equal)
	}
}

func EqualPairEq[A Eq[A], B Eq[B]]() Equaler[*Pair[A, B]] {
	return EqualPairBy(Equal[A], Equal[B])
}

func EqualPair[A comparable, B comparable]() Equaler[*Pair[A, B]] {
	return EqualPairBy(builtin.Equal[A], builtin.Equal[B])
}

func EqualPairBy[A, B any](fst Equaler[A], snd Equaler[B]) Equaler[*Pair[A, B]] {
	return EqualBy[*Pair[A, B]](
		function.On(fst, Fst[A, B]),
		function.On(snd, Snd[A, B]))
}

func EqualBy[A any](equals ...Equaler[A]) Equaler[A] {
	return EqualBys(equals)
}

func EqualBys[A any](equals []Equaler[A]) Equaler[A] {
	return func(x A, y A) bool {
		return All(func(c Equaler[A]) bool {
			return c(x, y)
		}, equals)
	}
}
