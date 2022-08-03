package slice

import (
	. "github.com/mattfenwick/collections/pkg/base"
	"github.com/mattfenwick/collections/pkg/builtin"
)

// Cons prepends an element to a slice.  This implementation may be inefficient.
func Cons[A any](a A, xs []A) []A {
	return append([]A{a}, xs...)
}

func Range[A builtin.Number](start A, stop A, step A) []A {
	current := start
	var out []A
	for {
		if current >= stop {
			break
		}
		out = append(out, current)
		current += step
	}
	return out
}

func IndexEq[T Eq[T]](s []T, e T) int {
	return IndexBy(Equal[T], s, e)
}

func Index[T comparable](s []T, e T) int {
	return IndexBy(builtin.Equal[T], s, e)
}

func IndexBy[T any](equal Equaler[T], s []T, e T) int {
	for i, v := range s {
		if equal(e, v) {
			return i
		}
	}
	return -1
}

func IsPrefixOfEq[A Eq[A]](xs []A, ys []A) bool {
	return IsPrefixOfBy(Equal[A], xs, ys)
}

func IsPrefixOf[A comparable](xs []A, ys []A) bool {
	return IsPrefixOfBy(builtin.Equal[A], xs, ys)
}

func IsPrefixOfBy[A any](equal Equaler[A], xs []A, ys []A) bool {
	for i := 0; i < len(xs); i++ {
		if i >= len(ys) || !equal(xs[i], ys[i]) {
			return false
		}
	}
	return true
}

func CartesianProduct[A, B any](xs []A, ys []B) []*Pair[A, B] {
	return CartesianProductWith(NewPair[A, B], xs, ys)
}

func CartesianProductWith[A, B, C any](f func(A, B) C, xs []A, ys []B) []C {
	var out []C
	for _, x := range xs {
		for _, y := range ys {
			out = append(out, f(x, y))
		}
	}
	return out
}
