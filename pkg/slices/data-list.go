package slices

import (
	"github.com/mattfenwick/collections/pkg"
	. "github.com/mattfenwick/collections/pkg/base"
	"github.com/mattfenwick/collections/pkg/builtins"
	"github.com/mattfenwick/collections/pkg/functions"
)

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
func Uncons[A any](xs []A) *pkg.Maybe[*Pair[A, []A]] {
	if len(xs) == 0 {
		return pkg.Nothing[*Pair[A, []A]]()
	}
	return pkg.Just[*Pair[A, []A]](NewPair(xs[0], xs[1:]))
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
func Map[A, B any](f F1[A, B], xs []A) []B {
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
func Foldl[A, B any](combine F2[B, A, B], base B, xs []A) B {
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
func Foldr[A, B any](combine F2[A, B, B], base B, xs []A) B {
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
func ConcatMap[A any, B any](f F1[A, []B], xs []A) []B {
	return Concat(Map(f, xs))
}

// And is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:and
func And(xs []bool) bool {
	return Foldl(builtins.And, true, xs)
}

// Or is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:or
func Or(xs []bool) bool {
	return Foldl(builtins.Or, false, xs)
}

// Any is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:any
func Any[A any](f F1[A, bool], xs []A) bool {
	return Foldl(builtins.Or, false, Map(f, xs))
}

// All is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:all
func All[A any](f F1[A, bool], xs []A) bool {
	return Foldl(builtins.And, true, Map(f, xs))
}

// Sum is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:sum
func Sum[A builtins.Number](xs []A) A {
	return Foldl[A, A](builtins.Plus[A], 0, xs)
}

// Product is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:product
func Product[A builtins.Number](xs []A) A {
	return Foldl[A, A](builtins.Times[A], 1, xs)
}

// Maximum is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:maximum
//   since Haskell's maximum blows up on empty lists, this has been modified for safety
func Maximum[A Ord[A]](xs []A) *pkg.Maybe[A] {
	if len(xs) == 0 {
		return pkg.Nothing[A]()
	}
	max := xs[0]
	for _, x := range xs[1:] {
		if GreaterThan(x, max) {
			max = x
		}
	}
	return pkg.Just(max)
}

// Minimum is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:minimum
//   since Haskell's minimum blows up on empty lists, this has been modified for safety
func Minimum[A Ord[A]](xs []A) *pkg.Maybe[A] {
	if len(xs) == 0 {
		return pkg.Nothing[A]()
	}
	min := xs[0]
	for _, x := range xs[1:] {
		if LessThan(x, min) {
			min = x
		}
	}
	return pkg.Just(min)
}

// Scanl is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:scanl
func Scanl[A, B any](combine F2[B, A, B], base B, xs []A) []B {
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
func Scanl1[A any](combine F2[A, A, A], xs []A) []A {
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
func Scanr[A, B any](combine F2[A, B, B], base B, xs []A) []B {
	state := base
	out := []B{base}
	for i := len(xs) - 1; i >= 0; i-- {
		state = combine(xs[i], state)
		out = Cons(state, out)
	}
	return out
}

// Scanr1 is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:scanr1
func Scanr1[A any](combine F2[A, A, A], xs []A) []A {
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

// MapAccumL is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:mapAccumL
//   forall t s a b. Traversable t => (s -> a -> (s, b)) -> s -> t a -> (s, t b)
func MapAccumL[A, B, S any](accum F2[S, A, *Pair[S, B]], s S, xs []A) *Pair[S, []B] {
	return Foldl(
		func(p *Pair[S, []B], a A) *Pair[S, []B] {
			p2 := accum(p.Fst, a)
			return NewPair[S, []B](p2.Fst, Append(p.Snd, []B{p2.Snd}))
		},
		NewPair[S, []B](s, nil),
		xs)
}

// MapAccumR is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:mapAccumR
//   forall t s a b. Traversable t => (s -> a -> (s, b)) -> s -> t a -> (s, t b)
func MapAccumR[A, B, S any](accum F2[S, A, *Pair[S, B]], s S, xs []A) *Pair[S, []B] {
	return Foldr(
		func(a A, p *Pair[S, []B]) *Pair[S, []B] {
			p2 := accum(p.Fst, a)
			return NewPair[S, []B](p2.Fst, Cons(p2.Snd, p.Snd))
		},
		NewPair[S, []B](s, nil),
		xs)
}

// Iterate is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:iterate
//   it uses a count to avoid an infinite slice
func Iterate[A any](count int, f F1[A, A], start A) []A {
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
	return Iterate(count, functions.Id[A], a)
}

// cycle can not be implemented in slices: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:cycle
//   (infinite list)

// Unfoldr is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:unfoldr
func Unfoldr[A, B any](f F1[B, *pkg.Maybe[*Pair[A, B]]], b B) []A {
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

// Take is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:take
func Take[A any](count int, xs []A) []A {
	if count > len(xs) {
		return xs
	}
	return xs[:count]
}

// Drop is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:drop
func Drop[A any](count int, xs []A) []A {
	if count > len(xs) {
		return nil
	}
	return xs[count:]
}

// SplitAt is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:splitAt
func SplitAt[A any](count int, xs []A) *Pair[[]A, []A] {
	return NewPair(Take(count, xs), Drop(count, xs))
}

// TakeWhile is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:takeWhile
func TakeWhile[A any](pred F1[A, bool], xs []A) []A {
	var out []A
	for _, x := range xs {
		if !pred(x) {
			break
		}
		out = append(out, x)
	}
	return out
}

// DropWhile is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:dropWhile
func DropWhile[A any](pred F1[A, bool], xs []A) []A {
	for i, x := range xs {
		if !pred(x) {
			return append([]A{}, xs[i:]...)
		}
	}
	return nil
}

// TODO
// DropWhileEnd is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:dropWhileEnd

// Span is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:span
func Span[A any](pred F1[A, bool], xs []A) *Pair[[]A, []A] {
	return NewPair(TakeWhile(pred, xs), DropWhile(pred, xs))
}

// Break is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:break
func Break[A any](pred F1[A, bool], xs []A) *Pair[[]A, []A] {
	// TODO are these type annotations all necessary?
	var f F1[A, bool] = functions.Compose(builtins.Not, pred)
	var g F1[F1[A, bool], F1[[]A, *Pair[[]A, []A]]] = functions.Partial2(Span[A])
	var h F1[[]A, *Pair[[]A, []A]] = g(f)
	return h(xs)
}

// TODO
// StripPrefix is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:stripPrefix

// TODO
// Group is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:group

// TODO
// Inits is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:inits

// TODO
// Tails is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:tails

// TODO Predicates: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#g:12

// TODO searching by equality

// Find is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:find
func Find[A any](pred F1[A, bool], xs []A) *pkg.Maybe[A] {
	for _, x := range xs {
		if pred(x) {
			return pkg.Just(x)
		}
	}
	return pkg.Nothing[A]()
}

// Filter is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:filter
func Filter[A any](f F1[A, bool], xs []A) []A {
	var out []A
	for _, x := range xs {
		if f(x) {
			out = append(out, x)
		}
	}
	return out
}

// Partition is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:partition
func Partition[A any](predicate F1[A, bool], xs []A) *Pair[[]A, []A] {
	var yes, no []A
	for _, x := range xs {
		if predicate(x) {
			yes = append(yes, x)
		} else {
			no = append(no, x)
		}
	}
	return NewPair(yes, no)
}

// TODO indexing lists: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#g:16

// Zip is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:zip
func Zip[A, B any](xs []A, ys []B) []*Pair[A, B] {
	return ZipWith(NewPair[A, B], xs, ys)
}

// TODO Zip3
// TODO Zip4
// TODO Zip5
// TODO Zip6
// TODO Zip7

// ZipWith is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:zipWith
func ZipWith[A, B, C any](f F2[A, B, C], xs []A, ys []B) []C {
	var out []C
	for i, x := range xs {
		if i >= len(ys) {
			break
		}
		out = append(out, f(x, ys[i]))
	}
	return out
}

// TODO ZipWith3
// TODO ZipWith4
// TODO ZipWith5
// TODO ZipWith6
// TODO ZipWith7

// Unzip is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:unzip
func Unzip[A, B any](xs []*Pair[A, B]) *Pair[[]A, []B] {
	var as []A
	var bs []B
	for _, x := range xs {
		as = append(as, x.Fst)
		bs = append(bs, x.Snd)
	}
	return NewPair(as, bs)
}

// TODO Unzip3
// TODO Unzip4
// TODO Unzip5
// TODO Unzip6
// TODO Unzip7

// TODO functions on strings: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#g:19

// TODO set operations: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#g:20

// Sort is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:sort
//   It orders elements by their natural Ord instance.
func Sort[A Ord[A]](xs []A) []A {
	return MergeSortWithComparator(Compare[A], xs)
}

// SortOn is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:sortOn
//   It uses the decorate/sort/undecorate pattern.
//   It allows a projection of each element to be used to determine the order.
//   The projection must have an Ord instance.
func SortOn[A any, B Ord[B]](projection F1[A, B], xs []A) []A {
	return SortOnBy(projection, Compare[B], xs)
}

// Insert is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:insert
func Insert[A Ord[A]](a A, xs []A) []A {
	panic("TODO")
}

// TODO the by operations: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#g:23

// SortBy is from: https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html#v:sortBy
//   It allows sorting based on a custom comparison operator;
//   therefore it does not require input elements to have an Ord instance.
func SortBy[A any](xs []A, compare F2[A, A, Ordering]) []A {
	//return SortOnBy(xs, Id[A], f)
	return MergeSortWithComparator(compare, xs)
}
