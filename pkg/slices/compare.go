package slices

import (
	. "github.com/mattfenwick/collections/pkg/base"
	"github.com/mattfenwick/collections/pkg/builtins"
	"github.com/mattfenwick/collections/pkg/functions"
	"golang.org/x/exp/constraints"
)

func CompareSliceIndexOrd[A Ord[A]](i int) Comparator[[]A] {
	return CompareSliceIndexBy(i, Compare[A])
}

func CompareSliceIndexOrdered[A constraints.Ordered](i int) Comparator[[]A] {
	return CompareSliceIndexBy(i, builtins.CompareOrdered[A])
}

// CompareSliceIndexBy compares a single index
func CompareSliceIndexBy[A any](i int, compare Comparator[A]) Comparator[[]A] {
	return func(xs []A, ys []A) Ordering {
		if i < len(xs) && i < len(ys) {
			return compare(xs[i], ys[i])
		} else if i >= len(xs) && i >= len(ys) {
			return OrderingEqual
		} else if i >= len(xs) {
			// ran off the end of xs
			return OrderingGreaterThan
		} else {
			// ran off the end of ys
			return OrderingLessThan
		}
	}
}

func CompareSlicePairwiseOrd[A Ord[A]]() Comparator[[]A] {
	return CompareSlicePairwiseBy(Compare[A])
}

func CompareSlicePairwiseOrdered[A constraints.Ordered]() Comparator[[]A] {
	return CompareSlicePairwiseBy(builtins.CompareOrdered[A])
}

// CompareSlicePairwiseBy should work as in Haskell.  Examples from Haskell:
//   Prelude> [1,2,3] < [3,4,5]
//   True
//   Prelude> [1,2,3] < [3,4]
//   True
//   Prelude> [1,2,3] < []
//   False
func CompareSlicePairwiseBy[A any](compare Comparator[A]) Comparator[[]A] {
	return func(xs []A, ys []A) Ordering {
		i := 0
		for {
			if i == len(xs) && i == len(ys) {
				return OrderingEqual
			} else if i == len(xs) {
				return OrderingLessThan
			} else if i == len(ys) {
				return OrderingGreaterThan
			}
			comp := compare(xs[i], ys[i])
			if comp != OrderingEqual {
				return comp
			}
			i++
		}
	}
}

func ComparePairOrd[A Ord[A], B Ord[B]]() Comparator[*Pair[A, B]] {
	return ComparePairBy(Compare[A], Compare[B])
}

func ComparePairOrdered[A constraints.Ordered, B constraints.Ordered]() Comparator[*Pair[A, B]] {
	return ComparePairBy(builtins.CompareOrdered[A], builtins.CompareOrdered[B])
}

func ComparePairBy[A, B any](fst Comparator[A], snd Comparator[B]) Comparator[*Pair[A, B]] {
	return CompareBy[*Pair[A, B]](
		functions.On(fst, First[A, B]),
		functions.On(snd, Second[A, B]))
}

func CompareBy[A any](comparisons ...Comparator[A]) Comparator[A] {
	return CompareBys(comparisons)
}

func CompareBys[A any](comparisons []Comparator[A]) Comparator[A] {
	return func(x A, y A) Ordering {
		ords := Map(func(c Comparator[A]) Ordering {
			return c(x, y)
		}, comparisons)
		for _, o := range ords {
			if o != OrderingEqual {
				return o
			}
		}
		return OrderingEqual
	}
}
