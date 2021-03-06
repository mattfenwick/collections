package function

import (
	. "github.com/mattfenwick/collections/pkg/base"
)

func Partial2[A, B, Z any](f F2[A, B, Z]) F1[A, F1[B, Z]] {
	return func(x A) F1[B, Z] {
		return func(y B) Z {
			return f(x, y)
		}
	}
}

func Partial3[A, B, C, Z any](f F3[A, B, C, Z]) F1[A, F1[B, F1[C, Z]]] {
	return func(a A) F1[B, F1[C, Z]] {
		return func(b B) F1[C, Z] {
			return func(c C) Z {
				return f(a, b, c)
			}
		}
	}
}

func Partial4[A, B, C, D, Z any](f F4[A, B, C, D, Z]) F1[A, F1[B, F1[C, F1[D, Z]]]] {
	return func(a A) F1[B, F1[C, F1[D, Z]]] {
		return func(b B) F1[C, F1[D, Z]] {
			return func(c C) F1[D, Z] {
				return func(d D) Z {
					return f(a, b, c, d)
				}
			}
		}
	}
}
