package slices

import (
	. "github.com/mattfenwick/collections/pkg/base"
)

func EqualSliceHelper[A any](compare F2[A, A, bool], xs []A, ys []A) bool {
	// unfortunately, can't do:
	//   return slices.Equal(xs, ys)
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

func EqualSlice[A any](compare F2[A, A, bool]) F2[[]A, []A, bool] {
	return func(xs []A, ys []A) bool {
		return EqualSliceHelper(compare, xs, ys)
	}
}

// CompareSliceHelper should work as in Haskell.  Examples from Haskell:
//   Prelude> [1,2,3] < [3,4,5]
//   True
//   Prelude> [1,2,3] < [3,4]
//   True
//   Prelude> [1,2,3] < []
//   False
func CompareSliceHelper[A any](compare F2[A, A, Ordering], xs []A, ys []A) Ordering {
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
func CompareSlice[A any](compare F2[A, A, Ordering]) F2[[]A, []A, Ordering] {
	return func(xs []A, ys []A) Ordering {
		return CompareSliceHelper(compare, xs, ys)
	}
}
