package functions

import (
	. "github.com/mattfenwick/collections/pkg/base"
)

func Compose[A, B, C any](f F1[B, C], g F1[A, B]) F1[A, C] {
	return func(x A) C {
		return f(g(x))
	}
}

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

func Const[A, B any](x A) F1[B, A] {
	return func(y B) A {
		return x
	}
}

func Id[A any](x A) A {
	return x
}

func Apply[A, B any](f F1[A, B], x A) B {
	return f(x)
}

func Flip[A, B, C any](f F2[A, B, C]) F2[B, A, C] {
	return func(b B, a A) C {
		return f(a, b)
	}
}

func On[A, B, C any](combine F2[B, B, C], project F1[A, B], x A, y A) C {
	return combine(project(x), project(y))
}
