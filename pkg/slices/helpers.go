package slices

import "github.com/mattfenwick/collections/pkg"

// Cons prepends an element to a slice.  This implementation may be inefficient.
func Cons[A any](a A, xs []A) []A {
	return append([]A{a}, xs...)
}

func Range[A pkg.Number](start A, stop A, step A) []A {
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
