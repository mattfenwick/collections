package slices

import "github.com/mattfenwick/collections/pkg/functions"

func GroupComparable[A comparable](xs []A) map[A][]A {
	return GroupOn(functions.Id[A], xs)
}

func GroupOn[A any, B comparable](projection func(A) B, xs []A) map[B][]A {
	out := map[B][]A{}
	for _, x := range xs {
		key := projection(x)
		slice := out[key]
		out[key] = append(slice, x)
	}
	return out
}
