package builtins

import (
	"github.com/mattfenwick/collections/pkg/base"
	"golang.org/x/exp/constraints"
)

func Equal[T comparable](a T, b T) bool {
	return EQ(a, b)
}

func CompareOrdered[A constraints.Ordered](a A, b A) base.Ordering {
	if a < b {
		return base.OrderingLessThan
	} else if a == b {
		return base.OrderingEqual
	} else {
		return base.OrderingGreaterThan
	}
}

func CompareBool(a bool, b bool) base.Ordering {
	if a == b {
		return base.OrderingEqual
	} else if !a {
		return base.OrderingLessThan
	}
	return base.OrderingGreaterThan
}
