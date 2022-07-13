package builtins

import (
	"github.com/mattfenwick/collections/pkg"
	"golang.org/x/exp/constraints"
)

func Compare[A constraints.Ordered](a A, b A) pkg.Ordering {
	if a < b {
		return pkg.OrderingLessThan
	} else if a == b {
		return pkg.OrderingEqual
	} else {
		return pkg.OrderingGreaterThan
	}
}

func Comparing[A constraints.Ordered, B any](f pkg.F1[B, A], x B, y B) pkg.Ordering {
	return Compare(f(x), f(y))
}

// ComparingP is a partial application of Comparing, fixing the first argument
func ComparingP[A constraints.Ordered, B any](f pkg.F1[B, A]) pkg.F2[B, B, pkg.Ordering] {
	return func(x B, y B) pkg.Ordering {
		return Comparing(f, x, y)
	}
}

func Sort[A constraints.Ordered](xs []A) []A {
	return pkg.SortBy(xs, Compare[A])
}

func SortOn[A any, B constraints.Ordered](xs []A, f pkg.F1[A, B]) []A {
	return pkg.MergeSortWithComparator(xs, ComparingP(f))
}
