package dict

import (
	. "github.com/mattfenwick/collections/pkg/base"
	"github.com/mattfenwick/collections/pkg/slice"
	"golang.org/x/exp/constraints"
)

func CompareIndexOrd[A comparable, B Ord[B]](key A) Comparator[map[A]B] {
	return CompareIndexBy(key, Compare[B])
}

func CompareIndex[A comparable, B constraints.Ordered](key A) Comparator[map[A]B] {
	return CompareIndexBy(key, CompareOrdered[B])
}

// CompareIndexBy compares a single index
func CompareIndexBy[A comparable, B any](key A, compare Comparator[B]) Comparator[map[A]B] {
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

func ComparePairwiseOrd[A constraints.Ordered, B Ord[B]]() Comparator[map[A]B] {
	return ComparePairwiseBy[A, B](Compare[B])
}

func ComparePairwise[A constraints.Ordered, B constraints.Ordered]() Comparator[map[A]B] {
	return ComparePairwiseBy[A, B](CompareOrdered[B])
}

// ComparePairwiseBy works by projecting a map to a list, therefore it's inefficient and probably
//
//	best to avoid using unless absolutely necessary!
//	Note: while `map` requires `A` to be `comparable`, *comparing* maps requires `A` to be Ordered
//	as well!
func ComparePairwiseBy[A constraints.Ordered, B any](compare Comparator[B]) Comparator[map[A]B] {
	comparePair := ComparePairBy(CompareOrdered[A], compare)
	compareSlice := slice.ComparePairwiseBy(comparePair)
	return func(xs map[A]B, ys map[A]B) Ordering {
		return compareSlice(
			slice.SortBy(ComparePairBy(CompareOrdered[A], ConstComparator[B](OrderingEqual)), ToSlice(xs)),
			slice.SortBy(ComparePairBy(CompareOrdered[A], ConstComparator[B](OrderingEqual)), ToSlice(ys)))
	}
}
