package base

import "fmt"

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

func (a Bool) Compare(b Bool) Ordering {
	if a == b {
		return OrderingEqual
	}
	if !a {
		return OrderingLessThan
	}
	return OrderingGreaterThan
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
