package slices

import (
	"fmt"
	. "github.com/mattfenwick/collections/pkg/base"
	"github.com/mattfenwick/collections/pkg/builtins"
	"github.com/mattfenwick/collections/pkg/functions"
)

func EqualSliceHelper[A any](compare F2[A, A, bool], xs []A, ys []A) bool {
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
func CompareSliceHelper[A any](compare Comparator[A], xs []A, ys []A) Ordering {
	i := 0
	for {
		if i == len(xs) && i == len(ys) {
			fmt.Println("at end of both slices: Equal")
			return OrderingEqual
		} else if i == len(xs) {
			fmt.Println("at end of first slice: LT")
			return OrderingLessThan
		} else if i == len(ys) {
			fmt.Println("at end of second slice: GT")
			return OrderingGreaterThan
		}
		comp := compare(xs[i], ys[i])
		if comp != OrderingEqual {
			fmt.Printf("element %d is %s, %+v, %+v\n", i, comp, xs, ys)
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

// TODO why is this here?
func On[A, B, C any](combine F2[B, B, C], projections []F1[A, B]) F2[A, A, []C] {
	return func(x A, y A) []C {
		return Map(func(p F1[A, B]) C {
			return functions.OnHelper(combine, p, x, y)
		}, projections)
	}
}

func EqualPair[A Eq[A], B Eq[B]](p1 *Pair[A, B], p2 *Pair[A, B]) bool {
	return EqualSliceHelper(
		builtins.Equal[bool],
		[]bool{p1.Fst.Equal(p2.Fst), p1.Snd.Equal(p2.Snd)},
		Replicate[bool](2, true))
}

func ComparePair[A Ord[A], B Ord[B]](p1 *Pair[A, B], p2 *Pair[A, B]) Ordering {
	return OrderedComparatorSplat[*Pair[A, B]](
		functions.On(Compare[A], First[A, B]),
		functions.On(Compare[B], Second[A, B]))(p1, p2)
}

// TODO this seems useless?
func ComparePairBy[A, B any](comparisons ...Comparator[*Pair[A, B]]) Comparator[*Pair[A, B]] {
	return func(p1 *Pair[A, B], p2 *Pair[A, B]) Ordering {
		return OrderedComparator(comparisons)(p1, p2)
	}
}

func OrderedComparatorSplat[A any](comparisons ...Comparator[A]) Comparator[A] {
	return OrderedComparator(comparisons)
}

func OrderedComparator[A any](comparisons []Comparator[A]) Comparator[A] {
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
