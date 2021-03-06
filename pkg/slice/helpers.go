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
