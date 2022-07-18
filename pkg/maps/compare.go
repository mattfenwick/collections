package maps

import (
	. "github.com/mattfenwick/collections/pkg/base"
	"github.com/mattfenwick/collections/pkg/builtins"
	"github.com/mattfenwick/collections/pkg/slices"
	"golang.org/x/exp/constraints"
)

func CompareMapIndexOrd[A comparable, B Ord[B]](key A) Comparator[map[A]B] {
	return CompareMapIndexBy(key, Compare[B])
}

func CompareMapIndex[A comparable, B constraints.Ordered](key A) Comparator[map[A]B] {
	return CompareMapIndexBy(key, builtins.CompareOrdered[B])
}

// CompareMapIndexBy compares a single index
func CompareMapIndexBy[A comparable, B any](key A, compare Comparator[B]) Comparator[map[A]B] {
	return func(xs map[A]B, ys map[A]B) Ordering {
		x, xok := xs[key]
		y, yok := ys[key]
		if xok && yok {
			return compare(x, y)
		} else if !xok && !yok {
			return OrderingEqual
		} else if !xok {
			// key not found in x
			return OrderingLessThan
		} else {
			// key not found in y
			return OrderingGreaterThan
		}
	}
}

func CompareMapPairwiseOrd[A constraints.Ordered, B Ord[B]]() Comparator[map[A]B] {
	return CompareMapPairwiseBy[A, B](Compare[B])
}

func CompareMapPairwise[A constraints.Ordered, B constraints.Ordered]() Comparator[map[A]B] {
	return CompareMapPairwiseBy[A, B](builtins.CompareOrdered[B])
}

// CompareMapPairwiseBy works by project a map to a list, therefore it's inefficient and probably
//   best to avoid using unless absolutely necessary!
//   Note: while `map` requires `A` to be `comparable`, *comparing* maps requires `A` to be Ordered
//   as well!
func CompareMapPairwiseBy[A constraints.Ordered, B any](compare Comparator[B]) Comparator[map[A]B] {
	comparePair := slices.ComparePairBy(builtins.CompareOrdered[A], compare)
	compareSlice := slices.CompareSlicePairwiseBy(comparePair)
	return func(xs map[A]B, ys map[A]B) Ordering {
		return compareSlice(
			slices.SortBy(slices.ComparePairBy(builtins.CompareOrdered[A], ConstComparator[B](OrderingEqual)), ToSlice(xs)),
			slices.SortBy(slices.ComparePairBy(builtins.CompareOrdered[A], ConstComparator[B](OrderingEqual)), ToSlice(ys)))
	}
}
