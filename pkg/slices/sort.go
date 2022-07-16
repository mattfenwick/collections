package slices

import (
	. "github.com/mattfenwick/collections/pkg/base"
)

// SortOnBy combines the functionality of `SortOn` and `SortBy`,
//   thereby separating projection and comparison functions
func SortOnBy[A any, B any](projection F1[A, B], compare F2[B, B, Ordering], xs []A) []A {
	pairs := Map(func(a A) *Pair[A, B] { return NewPair(a, projection(a)) }, xs)
	sorted := MergeSortWithComparator(func(p1 *Pair[A, B], p2 *Pair[A, B]) Ordering {
		return compare(p1.Snd, p2.Snd)
	}, pairs)
	return Map(First[A, B], sorted)
}

// MergeSortWithComparator needs to be rewritten iteratively TODO
func MergeSortWithComparator[A any](compare func(A, A) Ordering, xs []A) []A {
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
func Merge[A any](xs []A, ys []A, compare func(A, A) Ordering) []A {
	x, y := 0, 0
	var out []A
	for {
		if len(xs) == x {
			return append(out, ys[y:]...)
		} else if len(ys) == y {
			return append(out, xs[x:]...)
		}
		if compare(xs[x], ys[y]) == OrderingLessThan {
			out = append(out, xs[x])
			x++
		} else {
			out = append(out, ys[y])
			y++
		}
	}
}
