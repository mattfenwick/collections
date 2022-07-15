package pkg

import "golang.org/x/exp/constraints"

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

func And(a bool, b bool) bool {
	return a && b
}

func Or(a bool, b bool) bool {
	return a || b
}

func Not(a bool) bool {
	return !a
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
