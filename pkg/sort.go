package pkg

func (xs SliceOrd[A]) Sort() SliceOrd[A] {
	return Sort(xs)
}

// Sort orders elements by their natural Ord instance.
func Sort[A Ord[A]](xs []A) []A {
	return MergeSortWithComparator(xs, Compare[A])
}

// SortBy allows sorting based on a custom comparison operator;
//   therefore it does not require input elements to have an Ord instance.
//   See https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:sortBy
func SortBy[A any](xs []A, compare F2[A, A, Ordering]) []A {
	//return SortOnBy(xs, Id[A], f)
	return MergeSortWithComparator(xs, compare)
}

// SortOn is based on a Haskell function and the decorate/sort/undecorate pattern.
//   It allows a projection of each element to be used to determine the order.
//   The projection must have an Ord instance.
//   https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:sortOn
func SortOn[A any, B Ord[B]](xs []A, projection F1[A, B]) []A {
	return SortOnBy(xs, projection, Compare[B])
}

// SortOnBy combines the functionality of `SortOn` and `SortBy`,
//   thereby separating projection and comparison functions
func SortOnBy[A any, B any](xs []A, projection F1[A, B], compare F2[B, B, Ordering]) []A {
	pairs := MapSlice(func(a A) *Pair[A, B] { return NewPair(a, projection(a)) }, xs)
	sorted := MergeSortWithComparator(pairs, func(p1 *Pair[A, B], p2 *Pair[A, B]) Ordering {
		return compare(p1.Snd, p2.Snd)
	})
	return MapSlice(First[A, B], sorted)
}

// MergeSortWithComparator needs to be rewritten iteratively TODO
func MergeSortWithComparator[A any](xs []A, compare func(A, A) Ordering) []A {
	switch len(xs) {
	case 0, 1:
		return xs
	default:
		middle := len(xs) / 2
		return Merge(
			MergeSortWithComparator(xs[:middle], compare),
			MergeSortWithComparator(xs[middle:], compare),
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
