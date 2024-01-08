package base

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

// see https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-Ord.html#v:comparing for ideas

type Ordering string

const (
	OrderingLessThan    = "OrderingLessThan"
	OrderingEqual       = "OrderingEqual"
	OrderingGreaterThan = "OrderingGreaterThan"
)

func (a Ordering) Equal(b Ordering) bool {
	return a == b
}

func (a Ordering) Compare(b Ordering) Ordering {
	if a < b {
		return OrderingLessThan
	} else if a == b {
		return OrderingEqual
	}
	return OrderingGreaterThan
}

func (a Ordering) Flip() Ordering {
	switch a {
	case OrderingLessThan:
		return OrderingGreaterThan
	case OrderingEqual:
		return OrderingEqual
	case OrderingGreaterThan:
		return OrderingLessThan
	default:
		panic(fmt.Sprintf("invalid Ordering value: %s", a))
	}
}

type Comparator[A any] func(A, A) Ordering

func ConstComparator[T any](a Ordering) Comparator[T] {
	return func(_ T, _ T) Ordering { return a }
}

type Ord[T any] interface {
	Eq[T]
	Compare(T) Ordering
}

func Compare[A Ord[A]](x A, y A) Ordering {
	return x.Compare(y)
}

func FlipOrdering(a Ordering) Ordering {
	return a.Flip()
}

func LessThan[T Ord[T]](a T, b T) bool {
	return a.Compare(b) == OrderingLessThan
}

func LessThanOrEqual[T Ord[T]](a T, b T) bool {
	return a.Compare(b) != OrderingGreaterThan
}

func GreaterThan[T Ord[T]](a T, b T) bool {
	return a.Compare(b) == OrderingGreaterThan
}

func GreaterThanOrEqual[T Ord[T]](a T, b T) bool {
	return a.Compare(b) != OrderingLessThan
}

func Max[T Ord[T]](a T, b T) T {
	if GreaterThan(a, b) {
		return a
	}
	return b
}

func Min[T Ord[T]](a T, b T) T {
	if LessThan(a, b) {
		return a
	}
	return b
}

func CompareReverse[A any](compare Comparator[A]) Comparator[A] {
	return func(x A, y A) Ordering {
		return FlipOrdering(compare(x, y))
	}
}

func CompareOnOrd[A any, B Ord[B]](on func(A) B) Comparator[A] {
	return func(l A, r A) Ordering {
		return Compare(on(l), on(r))
	}
}

func CompareOrdered[A constraints.Ordered](a A, b A) Ordering {
	if a < b {
		return OrderingLessThan
	} else if a == b {
		return OrderingEqual
	} else {
		return OrderingGreaterThan
	}
}

func CompareOn[A any, B constraints.Ordered](on func(A) B) Comparator[A] {
	return func(l A, r A) Ordering {
		return CompareOrdered(on(l), on(r))
	}
}

func CompareBool(a bool, b bool) Ordering {
	if a == b {
		return OrderingEqual
	} else if !a {
		return OrderingLessThan
	}
	return OrderingGreaterThan
}

func (a Bool) Compare(b Bool) Ordering {
	if a == b {
		return OrderingEqual
	}
	if !a {
		return OrderingLessThan
	}
	return OrderingGreaterThan
}

// ComparatorToCmp converts a Comparator to a `cmp` function needed for using
//
//	golang's package `golang.org/x/exp/slices` sort functionality
//	 cmp(a, b) should return a negative number when a < b, a positive number when
//	 a > b and zero when a == b.
func ComparatorToCmp[A any](comparator Comparator[A]) func(A, A) int {
	return func(a A, b A) int {
		switch comparator(a, b) {
		case OrderingLessThan:
			return -1
		case OrderingEqual:
			return 0
		default:
			return 1
		}
	}
}

// OrdBox allows any constraints.Ordered to be used as an Ord
type OrdBox[A constraints.Ordered] struct {
	Value A
}

func (o *OrdBox[A]) Equal(other *OrdBox[A]) bool {
	return o.Value == other.Value
}

func (o *OrdBox[A]) Compare(other *OrdBox[A]) Ordering {
	if o.Value < other.Value {
		return OrderingLessThan
	} else if o.Value == other.Value {
		return OrderingEqual
	}
	return OrderingGreaterThan
}

func BoxOrd[A constraints.Ordered](a A) *OrdBox[A] {
	return &OrdBox[A]{Value: a}
}

func UnboxOrd[A constraints.Ordered](v *OrdBox[A]) A {
	return v.Value
}

// TODO how to sort complex numbers?  Python doesn't seem to support this?
//   maybe it's not a good idea?
//func (a Complex64) Compare(b Complex64) Ordering {
//	real(a)
//	if a < b {
//		return OrderingLessThan
//	} else if a == b {
//		return OrderingEqual
//	} else {
//		return OrderingGreaterThan
//	}
//}
//
//func (a Complex128) Compare(b Complex128) Ordering {
//	if a < b {
//		return OrderingLessThan
//	} else if a == b {
//		return OrderingEqual
//	} else {
//		return OrderingGreaterThan
//	}
//}
