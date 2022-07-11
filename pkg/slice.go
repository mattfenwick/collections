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

// foldl
// foldr
// scanl
// scanr
// any
// all
// reverse
// iterate
// repeat
// replicate
// take
// drop
// zip
