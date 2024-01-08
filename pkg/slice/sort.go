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

// ComparatorToLess was useful for working with older versions of the golang.org/x/exp/slices
//
//	package, whose sort functions used a different function type for comparison than they
//	currently do.  This function probably isn't useful anymore.
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
	slices.SortStableFunc(pairs, ComparatorToCmp(function.On(compare, Snd[A, B])))
	return Map(Fst[A, B], pairs)
}
