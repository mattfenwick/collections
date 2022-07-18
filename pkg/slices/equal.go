package slices

import (
	. "github.com/mattfenwick/collections/pkg/base"
	"github.com/mattfenwick/collections/pkg/builtins"
	"github.com/mattfenwick/collections/pkg/functions"
	"golang.org/x/exp/slices"
)

func EqualSliceIndexEq[A Eq[A]](i int) Equaler[[]A] {
	return EqualSliceIndexBy(i, Equal[A])
}

func EqualSliceIndexComparable[A comparable](i int) Equaler[[]A] {
	return EqualSliceIndexBy(i, builtins.Equal[A])
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

func EqualSlicePairwiseComparable[A comparable]() Equaler[[]A] {
	return EqualSlicePairwiseBy(builtins.Equal[A])
}

func EqualSlicePairwiseBy[A any](equal Equaler[A]) Equaler[[]A] {
	return func(xs []A, ys []A) bool {
		return slices.EqualFunc(xs, ys, equal)
	}
}

func EqualPairEq[A Eq[A], B Eq[B]]() Equaler[*Pair[A, B]] {
	return EqualPairBy(Equal[A], Equal[B])
}

func EqualPairComparable[A comparable, B comparable]() Equaler[*Pair[A, B]] {
	return EqualPairBy(builtins.Equal[A], builtins.Equal[B])
}

func EqualPairBy[A, B any](fst Equaler[A], snd Equaler[B]) Equaler[*Pair[A, B]] {
	return EqualBy[*Pair[A, B]](
		functions.On(fst, First[A, B]),
		functions.On(snd, Second[A, B]))
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
