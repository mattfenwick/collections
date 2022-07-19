package function

import (
	. "github.com/mattfenwick/collections/pkg/base"
)

// Id is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-Function.html#v:id
func Id[A any](x A) A {
	return x
}

// Const is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-Function.html#v:const
func Const[A, B any](x A) F1[B, A] {
	return func(y B) A {
		return x
	}
}

// Compose is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-Function.html#v:.
func Compose[A, B, C any](f F1[B, C], g F1[A, B]) F1[A, C] {
	return func(x A) C {
		return f(g(x))
	}
}

// Flip is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-Function.html#v:flip
func Flip[A, B, C any](f F2[A, B, C]) F2[B, A, C] {
	return func(b B, a A) C {
		return f(a, b)
	}
}

// Apply is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-Function.html#v:-36-
func Apply[A, B any](f F1[A, B], x A) B {
	return f(x)
}

// FlipApply is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-Function.html#v:-38-
func FlipApply[A, B any](x A, f F1[A, B]) B {
	return f(x)
}

// Fix is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-Function.html#v:fix
//   It doesn't seem to be possible in Golang so we'll leave it out.

// On is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-Function.html#v:on
//   It's partially applied because the expected use case is to specify the combine/project
//   args separately from the x/y args.
func On[A, B, C any](combine func(B, B) C, project func(A) B) func(A, A) C {
	return func(x A, y A) C {
		return OnHelper(combine, project, x, y)
	}
}
