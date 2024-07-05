package base

import "reflect"

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

// EqBoxBy allows any type to be used as an Eq
type EqBoxBy[A any] struct {
	Value A
	Eq    Equaler[A]
}

func (e *EqBoxBy[A]) Equal(other *EqBoxBy[A]) bool {
	return e.Eq(e.Value, other.Value)
}

func BoxEqBy[A any](a A, eq Equaler[A]) *EqBoxBy[A] {
	return &EqBoxBy[A]{Value: a, Eq: eq}
}

func UnboxEqBy[A any](v *EqBoxBy[A]) A {
	return v.Value
}

func EqualOnEq[A any, B Eq[B]](on func(A) B) Equaler[A] {
	return EqualOnBy(on, Equal[B])
}

func EqualOn[A any, B comparable](on func(A) B) Equaler[A] {
	return EqualOnBy(on, EqualComparable[B])
}

func EqualOnBy[A any, B any](on func(A) B, by Equaler[B]) Equaler[A] {
	return func(l A, r A) bool {
		return by(on(l), on(r))
	}
}

func ReflectDeepEqualer[A any](a A, b A) bool {
	return reflect.DeepEqual(a, b)
}
