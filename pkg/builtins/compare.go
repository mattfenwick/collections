package builtins

import (
	"github.com/mattfenwick/collections/pkg/base"
	"golang.org/x/exp/constraints"
)

func Compare[A constraints.Ordered](a A, b A) base.Ordering {
	if a < b {
		return base.OrderingLessThan
	} else if a == b {
		return base.OrderingEqual
	} else {
		return base.OrderingGreaterThan
	}
}

func Comparing[A constraints.Ordered, B any](f base.F1[B, A], x B, y B) base.Ordering {
	return Compare(f(x), f(y))
}

// ComparingP is a partial application of Comparing, fixing the first argument
func ComparingP[A constraints.Ordered, B any](f base.F1[B, A]) base.F2[B, B, base.Ordering] {
	return func(x B, y B) base.Ordering {
		return Comparing(f, x, y)
	}
}

// TODO are these necessary?
//func Sort[A constraints.Ordered](xs []A) []A {
//	return slices.SortBy(xs, Compare[A])
//}
//
//func SortOn[A any, B constraints.Ordered](xs []A, f base.F1[A, B]) []A {
//	return slices.MergeSortWithComparator(xs, ComparingP(f))
//}
