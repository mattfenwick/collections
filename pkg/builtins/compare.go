package builtins

import (
	"github.com/mattfenwick/collections/pkg/base"
	"github.com/mattfenwick/collections/pkg/functions"
	"golang.org/x/exp/constraints"
)

func Equal[T comparable](a T, b T) bool {
	return EQ(a, b)
}

func Compare[A constraints.Ordered](a A, b A) base.Ordering {
	if a < b {
		return base.OrderingLessThan
	} else if a == b {
		return base.OrderingEqual
	} else {
		return base.OrderingGreaterThan
	}
}

// TODO this should probably be deleted, it mixes the wrong things
func Comparing[A any, B constraints.Ordered](f base.F1[A, B], x A, y A) base.Ordering {
	return functions.On(Compare[B], f, x, y)
}

// TODO are these necessary?
//func Sort[A constraints.Ordered](xs []A) []A {
//	return slices.SortBy(xs, Compare[A])
//}
//
//func SortOn[A any, B constraints.Ordered](f base.F1[A, B], xs []A) []A {
//	return slices.SortOnBy(f, Compare[B], xs)
//}
