package base

// this example is from: https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#using-types-that-refer-to-themselves-in-constraints

type Equaler[A any] func(A, A) bool

type Eq[T any] interface {
	Equal(T) bool
}

func Equal[T Eq[T]](a T, b T) bool {
	return a.Equal(b)
}

func EqualComparable[T comparable](a T, b T) bool {
	return a == b
}

func NotEqual[T Eq[T]](a T, b T) bool {
	return !a.Equal(b)
}

func NotEqualComparable[T comparable](a T, b T) bool {
	return a != b
}

// EqBox allows any comparable to be used as an Eq
type EqBox[A comparable] struct {
	Value A
}

func (e *EqBox[A]) Equal(other *EqBox[A]) bool {
	return e.Value == other.Value
}

func BoxEq[A comparable](a A) *EqBox[A] {
	return &EqBox[A]{Value: a}
}

func UnboxEq[A comparable](v *EqBox[A]) A {
	return v.Value
}
