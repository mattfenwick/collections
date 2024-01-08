package dict

import (
	. "github.com/mattfenwick/collections/pkg/base"
	"github.com/mattfenwick/collections/pkg/slice"
	"golang.org/x/exp/constraints"
)

func CompareMapIndexOrd[A comparable, B Ord[B]](key A) Comparator[map[A]B] {
	return CompareMapIndexBy(key, Compare[B])
}

func CompareMapIndex[A comparable, B constraints.Ordered](key A) Comparator[map[A]B] {
	return CompareMapIndexBy(key, CompareOrdered[B])
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
	return CompareMapPairwiseBy[A, B](CompareOrdered[B])
}

// CompareMapPairwiseBy works by projecting a map to a list, therefore it's inefficient and probably
//
//	best to avoid using unless absolutely necessary!
//	Note: while `map` requires `A` to be `comparable`, *comparing* maps requires `A` to be Ordered
//	as well!
func CompareMapPairwiseBy[A constraints.Ordered, B any](compare Comparator[B]) Comparator[map[A]B] {
	comparePair := ComparePairBy(CompareOrdered[A], compare)
	compareSlice := slice.CompareSlicePairwiseBy(comparePair)
	return func(xs map[A]B, ys map[A]B) Ordering {
		return compareSlice(
			slice.SortBy(ComparePairBy(CompareOrdered[A], ConstComparator[B](OrderingEqual)), ToSlice(xs)),
			slice.SortBy(ComparePairBy(CompareOrdered[A], ConstComparator[B](OrderingEqual)), ToSlice(ys)))
	}
}
