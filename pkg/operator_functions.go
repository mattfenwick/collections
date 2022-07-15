package pkg

import (
	"golang.org/x/exp/constraints"
)

// Number is built out of:
//   https://pkg.go.dev/golang.org/x/exp@v0.0.0-20220706164943-b4a6d9510983/constraints
type Number interface {
	constraints.Integer | constraints.Float
}

func Plus[T Number](a T, b T) T {
	return a + b
}

func Minus[T Number](a T, b T) T {
	return a - b
}

func Times[T Number](a T, b T) T {
	return a * b
}

func Divide[T Number](a T, b T) T {
	return a / b
}

func Modulo[T constraints.Integer](a T, b T) T {
	return a % b
}

func Negate[T Number](a T) T {
	return -a
}

func And(a bool, b bool) bool {
	return a && b
}

func Or(a bool, b bool) bool {
	return a || b
}

func Not(a bool) bool {
	return !a
}

func LT[T constraints.Ordered](a T, b T) bool {
	return a < b
}

func LTE[T constraints.Ordered](a T, b T) bool {
	return a <= b
}

func GT[T constraints.Ordered](a T, b T) bool {
	return a > b
}

func GTE[T constraints.Ordered](a T, b T) bool {
	return a >= b
}

func EQ[T comparable](a T, b T) bool {
	return a == b
}

func NE[T comparable](a T, b T) bool {
	return a != b
}

func Slice[A any](start uint, stop uint, xs []A) []A {
	return xs[start:stop]
}

func SliceFrom[A any](start uint, xs []A) []A {
	return xs[start:]
}

func SliceTo[A any](stop uint, xs []A) []A {
	return xs[:stop]
}

func Reference[A any](a A) *A {
	return &a
}

func Dereference[A any](a *A) A {
	return *a
}

// TODO see https://go.dev/ref/spec#Operators_and_punctuation for a full list of operators
//   don't do the ones that:
//     aren't functional, for example: '+=' or '++'
//     aren't for expressions, e.g.: '...' or ':='
//   *do* do: & | ^ << >> &^
