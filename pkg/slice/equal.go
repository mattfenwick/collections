package slice

import (
	. "github.com/mattfenwick/collections/pkg/base"
	"github.com/mattfenwick/collections/pkg/builtin"
	"golang.org/x/exp/slices"
)

func EqualIndexEq[A Eq[A]](i int) Equaler[[]A] {
	return EqualIndexBy(i, Equal[A])
}

func EqualIndex[A comparable](i int) Equaler[[]A] {
	return EqualIndexBy(i, builtin.EQ[A])
}

// EqualIndexBy looks at a single index
func EqualIndexBy[A any](i int, equal Equaler[A]) Equaler[[]A] {
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

func EqualPairwiseEq[A Eq[A]]() Equaler[[]A] {
	return EqualPairwiseBy(Equal[A])
}

func EqualPairwise[A comparable]() Equaler[[]A] {
	return EqualPairwiseBy(builtin.EQ[A])
}

func EqualPairwiseBy[A any](equal Equaler[A]) Equaler[[]A] {
	return func(xs []A, ys []A) bool {
		return slices.EqualFunc(xs, ys, equal)
	}
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
