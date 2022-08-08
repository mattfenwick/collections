package slice

import (
	. "github.com/mattfenwick/collections/pkg/base"
	"golang.org/x/exp/slices"
)

func SortOrd[A Ord[A]](xs []A) []A {
	return SortBy(Compare[A], xs)
}

func SortOnOrd[A any, B Ord[B]](projection F1[A, B], xs []A) []A {
	return SortOnBy(projection, Compare[B], xs)
}

func ComparatorToLess[A any](comparator Comparator[A]) func(A, A) bool {
	return func(a A, b A) bool {
		return comparator(a, b) == OrderingLessThan
	}
}

// SortOnBy combines the functionality of `SortOn` and `SortBy`,
//
//	thereby separating projection and comparison functions
func SortOnBy[A any, B any](projection F1[A, B], compare Comparator[B], xs []A) []A {
	pairs := Map(func(a A) *Pair[A, B] { return NewPair(a, projection(a)) }, xs)
	slices.SortStableFunc(pairs, func(p1, p2 *Pair[A, B]) bool {
		return compare(p1.Snd, p2.Snd) == OrderingLessThan
	})
	return Map(Fst[A, B], pairs)
}
