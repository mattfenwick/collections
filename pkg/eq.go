package pkg

// this example is from: https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#using-types-that-refer-to-themselves-in-constraints

type Eq[T any] interface {
	Equal(T) bool
}

func NotEqual[T Eq[T]](a T, b T) bool {
	return !a.Equal(b)
}

func (xs SliceEq[A]) Equal(ys SliceEq[A]) bool {
	// unfortunately, can't do:
	//   return slices.Equal(xs, ys)
	//   because: A does not implement comparable
	if len(xs) != len(ys) {
		return false
	}
	for i := range xs {
		if !xs[i].Equal(ys[i]) {
			return false
		}
	}
	return true
}

func (xs MapEq[A, B]) Equal(ys MapEq[A, B]) bool {
	// unfortunately, can't do:
	//   return maps.Equal(xs, ys)
	//   because: B does not implement comparable
	if len(xs) != len(ys) {
		return false
	}
	for k, vx := range xs {
		vy, ok := ys[k]
		if !ok {
			return false
		}
		if !vx.Equal(vy) {
			return false
		}
	}
	return true
}

// TODO any way to use this?
//func EqualComparable[T comparable](a T, b T) bool {
//	return a == b
//}

func Index[T Eq[T]](s []T, e T) int {
	for i, v := range s {
		if e.Equal(v) {
			return i
		}
	}
	return -1
}
