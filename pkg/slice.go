package pkg

// for golang examples, see:
//   https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#map_reduce_filter for inspiration

// these functions are inspired by Haskell's Data.List:
//   https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html

func MapSlice[A, B any](f F1[A, B], xs []A) []B {
	out := make([]B, len(xs))
	for i, x := range xs {
		out[i] = f(x)
	}
	return out
}

func ReduceSlice[A, B any](f F2[B, A, B], b B, xs []A) B {
	out := b
	for _, x := range xs {
		b = f(b, x)
	}
	return out
}

func FilterSlice[A any](f F1[A, bool], xs []A) []A {
	var out []A
	for _, x := range xs {
		if f(x) {
			out = append(out, x)
		}
	}
	return out
}

func ReverseSlice[A any](xs []A) []A {
	length := len(xs)
	out := make([]A, length)
	for i, x := range xs {
		out[length-i-1] = x
	}
	return out
}

func AllSlice[A any](f F1[A, bool], xs []A) bool {
	return ReduceSlice[bool, bool](And, true, MapSlice(f, xs))
}

func AnySlice[A any](f F1[A, bool], xs []A) bool {
	return ReduceSlice(Or, false, MapSlice(f, xs))
}

// Replicate should be written using unfold or something
func Replicate[A any](count int, item A) []A {
	out := make([]A, count)
	for i := 0; i < count; i++ {
		out[i] = item
	}
	return out
}

func Zip[A, B any](xs []A, ys []B) []*Pair[A, B] {
	var out []*Pair[A, B]
	for i, x := range xs {
		if i >= len(ys) {
			break
		}
		out = append(out, NewPair[A, B](x, ys[i]))
	}
	return out
}

func ConcatSlice[A any](xs []A, ys []A) []A {
	return append(xs, ys...)
}

func ConcatSlices[A any](xss [][]A) []A {
	var out []A
	for _, xs := range xss {
		out = append(out, xs...)
	}
	return out
}

func IsEmptySlice[A any](xs []A) bool {
	return len(xs) == 0
}

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

func Intercalate[A any](sep []A, xss [][]A) []A {
	return ConcatSlices(Intersperse(sep, xss))
}

func SumSlice[A Number](xs []A) A {
	return ReduceSlice[A, A](Plus[A], 0, xs)
}

func ProductSlice[A Number](xs []A) A {
	return ReduceSlice[A, A](Times[A], 1, xs)
}

func UnfoldrSlice[A, B any](f F1[B, *Maybe[*Pair[A, B]]], b B) []A {
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

// foldl
// foldr
// scanl
// scanr
// iterate
// repeat
// take
// drop
// unzip
// unfold
