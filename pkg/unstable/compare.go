package unstable

import (
	"fmt"
	"github.com/mattfenwick/collections/pkg/base"
	"golang.org/x/exp/constraints"
)

// see https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-Ord.html#v:comparing for ideas

type Equaler[A any] func(A, A) bool

type Eq[T any] interface {
	Equal(T) bool
}

func Equal[T Eq[T]](a T, b T) bool {
	return a.Equal(b)
}

func NotEqual[T Eq[T]](a T, b T) bool {
	return !a.Equal(b)
}

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

func Compare[A constraints.Ordered](a A, b A) Ordering {
	if a < b {
		return base.OrderingLessThan
	} else if a == b {
		return base.OrderingEqual
	} else {
		return base.OrderingGreaterThan
	}
}

func CompareOrd[A Ord[A]](x A, y A) Ordering {
	return x.Compare(y)
}

func CompareOn[A any, B constraints.Ordered](on func(A) B) Comparator[A] {
	return func(l A, r A) Ordering {
		return Compare(on(l), on(r))
	}
}

func CompareOnOrd[A any, B Ord[B]](on func(A) B) Comparator[A] {
	return func(l A, r A) Ordering {
		return CompareOrd(on(l), on(r))
	}
}

func CompareReverse[A any](compare Comparator[A]) Comparator[A] {
	return func(x A, y A) Ordering {
		return compare(x, y).Flip()
	}
}

// CompareBool is necessary because: bool is not part of constraints.Ordered
func CompareBool(a bool, b bool) base.Ordering {
	if a == b {
		return base.OrderingEqual
	} else if !a {
		return base.OrderingLessThan
	}
	return base.OrderingGreaterThan
}
