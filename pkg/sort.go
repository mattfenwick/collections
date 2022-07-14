package pkg

func (xs SliceOrd[A]) Sort() SliceOrd[A] {
	return Sort(xs)
}

func Sort[A Ord[A]](xs []A) []A {
	return MergeSortWithComparator(xs, Compare[A])
}

func SortBy[A any](xs []A, f F2[A, A, Ordering]) []A {
	return MergeSortWithComparator(xs, f)
}

// SortOn is based on a Haskell function and the decorate/sort/undecorate pattern:
//   example: https://hackage.haskell.org/package/base-4.16.2.0/docs/src/Data.OldList.html#sortOn
func SortOn[A any, B Ord[B]](xs []A, f F1[A, B]) []A {
	pairs := MapSlice(func(a A) *Pair[A, B] { return NewPair(a, f(a)) }, xs)
	sorted := MergeSortWithComparator(pairs, ComparingP(Second[A, B]))
	return MapSlice(First[A, B], sorted)
}

// MergeSortWithComparator needs to be rewritten iteratively TODO
func MergeSortWithComparator[A any](xs []A, f func(A, A) Ordering) []A {
	switch len(xs) {
	case 0, 1:
		return xs
	default:
		middle := len(xs) / 2
		return Merge(MergeSortWithComparator(xs[:middle], f), MergeSortWithComparator(xs[middle:], f), f)
	}
}

// Merge ...
func Merge[A any](xs []A, ys []A, f func(A, A) Ordering) []A {
	x, y := 0, 0
	var out []A
	for {
		if len(xs) == x {
			return append(out, ys[y:]...)
		} else if len(ys) == y {
			return append(out, xs[x:]...)
		}
		if f(xs[x], ys[y]) == OrderingLessThan {
			out = append(out, xs[x])
			x++
		} else {
			out = append(out, ys[y])
			y++
		}
	}
}
