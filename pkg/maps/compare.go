package maps

/*
TODO
func EqualMapHelper[A any](compare F2[A, A, bool], xs []A, ys []A) bool {
	// unfortunately, can't do:
	//   return Equal(xs, ys)
	//   because: A does not implement comparable
	if len(xs) != len(ys) {
		return false
	}
	for i := range xs {
		if !compare(xs[i], ys[i]) {
			return false
		}
	}
	return true
}

func EqualMap[A any](compare F2[A, A, bool]) F2[[]A, []A, bool] {
	return func(xs []A, ys []A) bool {
		return EqualSliceHelper(compare, xs, ys)
	}
}

func EqualPair[A Eq[A], B Eq[B]](p1 *Pair[A, B], p2 *Pair[A, B]) bool {
	return EqualBy[*Pair[A, B]](
		functions.On(Equal[A], First[A, B]),
		functions.On(Equal[B], Second[A, B]))(p1, p2)
}

func EqualBy[A any](comparisons ...Equaler[A]) Equaler[A] {
	return EqualBys(comparisons)
}

func EqualBys[A any](comparisons []Equaler[A]) Equaler[A] {
	return func(x A, y A) bool {
		return All(func(c Equaler[A]) bool {
			return c(x, y)
		}, comparisons)
	}
}

// CompareSliceHelper should work as in Haskell.  Examples from Haskell:
//   Prelude> [1,2,3] < [3,4,5]
//   True
//   Prelude> [1,2,3] < [3,4]
//   True
//   Prelude> [1,2,3] < []
//   False
func CompareSliceHelper[A any](compare Comparator[A], xs []A, ys []A) Ordering {
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
func CompareSlice[A any](compare Comparator[A]) Comparator[[]A] {
	return func(xs []A, ys []A) Ordering {
		return CompareSliceHelper(compare, xs, ys)
	}
}

func ComparePair[A Ord[A], B Ord[B]](p1 *Pair[A, B], p2 *Pair[A, B]) Ordering {
	return CompareBy[*Pair[A, B]](
		functions.On(Compare[A], First[A, B]),
		functions.On(Compare[B], Second[A, B]))(p1, p2)
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
*/
