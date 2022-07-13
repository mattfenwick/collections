package pkg

type Ordering string

const (
	OrderingLessThan    = "OrderingLessThan"
	OrderingEqual       = "OrderingEqual"
	OrderingGreaterThan = "OrderingGreaterThan"
)

type Ord[T any] interface {
	Eq[T]
	Compare(T) Ordering
}

func Compare[A Ord[A]](y A, z A) Ordering {
	return y.Compare(z)
}

func Comparing[A Ord[A], B any](f F1[B, A], x B, y B) Ordering {
	return f(x).Compare(f(y))
}

// ComparingP is a partial application of Comparing, fixing the first argument
func ComparingP[A Ord[A], B any](f F1[B, A]) F2[B, B, Ordering] {
	return func(x B, y B) Ordering {
		return f(x).Compare(f(y))
	}
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

// Compare should work like in Haskell.  Examples from Haskell:
//   Prelude> [1,2,3] < [3,4,5]
//   True
//   Prelude> [1,2,3] < [3,4]
//   True
//   Prelude> [1,2,3] < []
//   False
func (xs SliceOrd[A]) Compare(ys SliceOrd[A]) Ordering {
	i := 0
	for {
		if i == len(xs) {
			return OrderingLessThan
		}
		if i == len(ys) {
			return OrderingGreaterThan
		}
		comp := xs[i].Compare(ys[i])
		if comp != OrderingEqual {
			return comp
		}
		i++
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
