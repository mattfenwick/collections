package pkg

// see https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#map_reduce_filter for inspiration

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
	return len(FilterSlice(f, xs)) == len(xs)
}

func AnySlice[A any](f F1[A, bool], xs []A) bool {
	return len(FilterSlice(f, xs)) > 0
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
