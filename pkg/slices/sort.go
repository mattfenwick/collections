package slices

import (
	. "github.com/mattfenwick/collections/pkg/base"
)

// SortOnBy combines the functionality of `SortOn` and `SortBy`,
//   thereby separating projection and comparison functions
func SortOnBy[A any, B any](projection F1[A, B], compare Comparator[B], xs []A) []A {
	pairs := Map(func(a A) *Pair[A, B] { return NewPair(a, projection(a)) }, xs)
	sorted := MergeSortWithComparator(func(p1 *Pair[A, B], p2 *Pair[A, B]) Ordering {
		return compare(p1.Snd, p2.Snd)
	}, pairs)
	return Map(First[A, B], sorted)
}

// MergeSortWithComparator needs to be rewritten iteratively TODO
func MergeSortWithComparator[A any](compare Comparator[A], xs []A) []A {
	switch len(xs) {
	case 0, 1:
		return xs
	default:
		middle := len(xs) / 2
		return Merge(
			MergeSortWithComparator(compare, xs[:middle]),
			MergeSortWithComparator(compare, xs[middle:]),
			compare)
	}
}

// Merge ...
func Merge[A any](xs []A, ys []A, compare Comparator[A]) []A {
	x, y := 0, 0
	out := make([]A, len(xs)+len(ys))
	for i := 0; ; i++ {
		if len(xs) == x {
			for ; y < len(ys); i, y = i+1, y+1 {
				out[i] = ys[y]
			}
			return out
		} else if len(ys) == y {
			for ; x < len(xs); i, x = i+1, x+1 {
				out[i] = xs[x]
			}
			return out
		}
		if compare(xs[x], ys[y]) == OrderingLessThan {
			out[i] = xs[x]
			x++
		} else {
			out[i] = ys[y]
			y++
		}
	}
}
