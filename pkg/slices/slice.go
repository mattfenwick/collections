package slices

import "github.com/mattfenwick/collections/pkg"

// this code is based on Haskell's data.List:
//   see https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html

// Append is from (++):
//   https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:-43--43-
func Append[A any](xs []A, ys []A) []A {
	out := append([]A{}, xs...)
	return append(out, ys...)
}

// TODO what to do with questionable/partial functions ?
//func Head[A any](xs []A) A {
//
//}
//
//func Last[A any](xs []A) A {
//
//}
//
//func Tail[A any](xs []A) []A {
//
//}
//
//func Init[A any](xs []A) []A {
//
//}

// Uncons is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:uncons
func Uncons[A any](xs []A) *pkg.Maybe[*pkg.Pair[A, []A]] {
	if len(xs) == 0 {
		return pkg.Nothing[*pkg.Pair[A, []A]]()
	}
	return pkg.Just[*pkg.Pair[A, []A]](pkg.NewPair(xs[0], xs[1:]))
}

// Singleton is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:singleton
func Singleton[A any](a A) []A {
	return []A{a}
}

// Null is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:null
func Null[A any](xs []A) bool {
	return len(xs) == 0
}

// Length is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:length
func Length[A any](xs []A) int {
	return len(xs)
}

// Map is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:map
func Map[A, B any](f pkg.F1[A, B], xs []A) []B {
	out := make([]B, len(xs))
	for i, x := range xs {
		out[i] = f(x)
	}
	return out
}

// Reverse is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:reverse
func Reverse[A any](xs []A) []A {
	length := len(xs)
	out := make([]A, length)
	for i, x := range xs {
		out[length-i-1] = x
	}
	return out
}

// Intersperse is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:intersperse
func Intersperse[A any](sep A, xs []A) []A {
	if len(xs) <= 1 {
		return xs
	}
	out := []A{xs[0]}
	for i := 1; i < len(xs); i++ {
		out = append(out, sep, xs[i])
	}
	return out
}

// Intercalate is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:intercalate
func Intercalate[A any](sep []A, xss [][]A) []A {
	return Concat(Intersperse(sep, xss))
}

// TODO transpose
//   see https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:transpose

// TODO https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:subsequences

// TODO https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:permutations

// Foldl is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:foldl
//   foldl f z [x1, x2, ..., xn] == (...((z `f` x1) `f` x2) `f`...) `f` xn
func Foldl[A, B any](combine pkg.F2[B, A, B], base B, xs []A) B {
	out := base
	for _, x := range xs {
		out = combine(out, x)
	}
	return out
}

// TODO foldl' https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:foldl-39-

// TODO foldl1 https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:foldl1

// TODO foldl1' https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:foldl1-39-

// Foldr is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:foldr
//   foldr f z [x1, x2, ..., xn] == x1 `f` (x2 `f` ... (xn `f` z)...)
func Foldr[A, B any](combine pkg.F2[A, B, B], base B, xs []A) B {
	out := base
	for i := len(xs) - 1; i >= 0; i-- {
		out = combine(xs[i], out)
	}
	return out
}

// TODO foldr1 https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:foldr1

// Concat is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:concat
func Concat[A any](xss [][]A) []A {
	var out []A
	for _, xs := range xss {
		out = append(out, xs...)
	}
	return out
}

// ConcatMap is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:concatMap
func ConcatMap[A any, B any](f pkg.F1[A, []B], xs []A) []B {
	return Concat(Map(f, xs))
}

// And is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:and
func And(xs []bool) bool {
	return Foldl(pkg.And, true, xs)
}

// Or is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:or
func Or(xs []bool) bool {
	return Foldl(pkg.Or, false, xs)
}

// Any is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:any
func Any[A any](f pkg.F1[A, bool], xs []A) bool {
	return Foldl(pkg.Or, false, Map(f, xs))
}

// All is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:all
func All[A any](f pkg.F1[A, bool], xs []A) bool {
	return Foldl(pkg.And, true, Map(f, xs))
}

// Sum is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:sum
func Sum[A pkg.Number](xs []A) A {
	return Foldl[A, A](pkg.Plus[A], 0, xs)
}

// Product is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:product
func Product[A pkg.Number](xs []A) A {
	return Foldl[A, A](pkg.Times[A], 1, xs)
}

// Maximum is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:maximum
//   since Haskell's maximum blows up on empty lists, this has been modified for safety
func Maximum[A pkg.Ord[A]](xs []A) *pkg.Maybe[A] {
	if len(xs) == 0 {
		return pkg.Nothing[A]()
	}
	max := xs[0]
	for _, x := range xs[1:] {
		if pkg.GreaterThan(x, max) {
			max = x
		}
	}
	return pkg.Just(max)
}

// Minimum is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:minimum
//   since Haskell's minimum blows up on empty lists, this has been modified for safety
func Minimum[A pkg.Ord[A]](xs []A) *pkg.Maybe[A] {
	if len(xs) == 0 {
		return pkg.Nothing[A]()
	}
	min := xs[0]
	for _, x := range xs[1:] {
		if pkg.LessThan(x, min) {
			min = x
		}
	}
	return pkg.Just(min)
}

// Scanl is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:scanl
func Scanl[A, B any](combine pkg.F2[B, A, B], base B, xs []A) []B {
	out := []B{base}
	state := base
	for _, x := range xs {
		state = combine(state, x)
		out = append(out, state)
	}
	return out
}

// TODO scanl' https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:scanl-39-
//   since the point of this function is to be strictly evaluated, this is probably an unnecessary
//   variant in golang

// Scanl1 is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:scanl1
func Scanl1[A any](combine pkg.F2[A, A, A], xs []A) []A {
	if len(xs) == 0 {
		return xs
	}
	state := xs[0]
	out := []A{state}
	for _, x := range xs[1:] {
		state = combine(state, x)
		out = append(out, state)
	}
	return out
}

// Scanr is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:scanr
func Scanr[A, B any](combine pkg.F2[A, B, B], base B, xs []A) []B {
	state := base
	out := []B{base}
	for i := len(xs) - 1; i >= 0; i-- {
		state = combine(xs[i], state)
		out = Cons(state, out)
	}
	return out
}

// Scanr1 is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:scanr1
func Scanr1[A any](combine pkg.F2[A, A, A], xs []A) []A {
	if len(xs) == 0 {
		return xs
	}
	state := xs[len(xs)-1]
	out := []A{state}
	for i := len(xs) - 2; i >= 0; i-- {
		state = combine(xs[i], state)
		out = Cons(state, out)
	}
	return out
}

// TODO
// MapAccumL is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:mapAccumL

// TODO
// MapAccumR is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:mapAccumR

// Iterate is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:iterate
//   it uses a count to avoid an infinite slice
func Iterate[A any](count int, f pkg.F1[A, A], start A) []A {
	if count == 0 {
		return nil
	}
	state := start
	out := []A{state}
	for i := 1; i < count; i++ {
		state = f(state)
		out = append(out, state)
	}
	return out
}

// iterate' is not necessary: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:iterate-39-
//   (strict version of iterate)

// repeat can not be implemented in slices: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:repeat
//   (infinite list)

// Replicate is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:replicate
func Replicate[A any](count int, a A) []A {
	return Iterate(count, pkg.Id[A], a)
}

// cycle can not be implemented in slices: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:cycle
//   (infinite list)

// Unfoldr is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:unfoldr
func Unfoldr[A, B any](f pkg.F1[B, *pkg.Maybe[*pkg.Pair[A, B]]], b B) []A {
	var out []A
	nextB := b
	for {
		next := f(nextB)
		if next.Value == nil {
			break
		}
		out = append(out, (*next.Value).Fst)
		nextB = (*next.Value).Snd
	}
	return out
}
