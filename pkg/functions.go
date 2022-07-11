package pkg

type F1[A, Z any] func(A) Z
type F2[A, B, Z any] func(A, B) Z
type F3[A, B, C, Z any] func(A, B, C) Z
type F4[A, B, C, D, Z any] func(A, B, C, D) Z
type F5[A, B, C, D, E, Z any] func(A, B, C, D, E) Z

func Compose[A, B, C any](f F1[B, C], g F1[A, B]) F1[A, C] {
	return func(x A) C {
		return f(g(x))
	}
}

func Partial2[A, B, C any](f F2[A, B, C]) F1[A, F1[B, C]] {
	return func(x A) F1[B, C] {
		return func(y B) C {
			return f(x, y)
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

func Replicate[A any](count int, item A) []A {
	out := make([]A, count)
	for i := 0; i < count; i++ {
		out[i] = item
	}
	return out
}

func Apply[A, B any](f F1[A, B], x A) B {
	return f(x)
}

func Flip[A, B, C any](f F2[A, B, C]) F2[B, A, C] {
	return func(b B, a A) C {
		return f(a, b)
	}
}
