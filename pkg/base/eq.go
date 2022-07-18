package base

// this example is from: https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#using-types-that-refer-to-themselves-in-constraints

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

// TODO any way to use this?
//func EqualComparable[T comparable](a T, b T) bool {
//	return a == b
//}
