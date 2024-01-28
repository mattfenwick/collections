package slice

import (
	. "github.com/mattfenwick/collections/pkg/base"
	"golang.org/x/exp/constraints"
)

func CompareIndexOrd[A Ord[A]](i int) Comparator[[]A] {
	return CompareIndexBy(i, Compare[A])
}

func CompareIndex[A constraints.Ordered](i int) Comparator[[]A] {
	return CompareIndexBy(i, CompareOrdered[A])
}

// CompareIndexBy compares a single index
func CompareIndexBy[A any](i int, compare Comparator[A]) Comparator[[]A] {
	return func(xs []A, ys []A) Ordering {
		if i < len(xs) && i < len(ys) {
			return compare(xs[i], ys[i])
		} else if i >= len(xs) && i >= len(ys) {
			return OrderingEqual
		} else if i >= len(xs) {
			// ran off the end of xs
			return OrderingLessThan
		} else {
			// ran off the end of ys
			return OrderingGreaterThan
		}
	}
}

func ComparePairwiseOrd[A Ord[A]]() Comparator[[]A] {
	return ComparePairwiseBy(Compare[A])
}

func ComparePairwise[A constraints.Ordered]() Comparator[[]A] {
	return ComparePairwiseBy(CompareOrdered[A])
}

// ComparePairwiseBy should work as in Haskell.  Examples from Haskell:
// Prelude> [1,2,3] < [3,4,5]
// True
// Prelude> [1,2,3] < [3,4]
// True
// Prelude> [1,2,3] < []
// False
func ComparePairwiseBy[A any](compare Comparator[A]) Comparator[[]A] {
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
