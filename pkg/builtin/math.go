package builtin

import "golang.org/x/exp/constraints"

func Max[T constraints.Ordered](a T, b T) T {
	if GT(a, b) {
		return a
	}
	return b
}

func Min[T constraints.Ordered](a T, b T) T {
	if LT(a, b) {
		return a
	}
	return b
}
