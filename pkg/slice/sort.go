package slice

import (
	. "github.com/mattfenwick/collections/pkg/base"
	"github.com/mattfenwick/collections/pkg/function"
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

// ComparatorToCmp converts a Comparator to a `cmp` function needed for using
//
//	golang's package `golang.org/x/exp/slices` sort functionality
//	 cmp(a, b) should return a negative number when a < b, a positive number when
//	 a > b and zero when a == b.
func ComparatorToCmp[A any](comparator Comparator[A]) func(A, A) int {
	return func(a A, b A) int {
		switch comparator(a, b) {
		case OrderingLessThan:
			return -1
		case OrderingEqual:
			return 0
		default:
			return 1
		}
	}
}

// SortOnBy combines the functionality of `SortOn` and `SortBy`,
//
//	thereby separating projection and comparison functions
func SortOnBy[A any, B any](projection F1[A, B], compare Comparator[B], xs []A) []A {
	pairs := Map(func(a A) *Pair[A, B] { return NewPair(a, projection(a)) }, xs)
	slices.SortStableFunc(pairs, ComparatorToCmp(function.On(compare, Snd[A, B])))
	return Map(Fst[A, B], pairs)
}
