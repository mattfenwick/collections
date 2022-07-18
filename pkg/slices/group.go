package slices

func GroupOn[A any, B comparable](projection func(A) B, xs []A) map[B][]A {
	out := map[B][]A{}
	for _, x := range xs {
		key := projection(x)
		slice := out[key]
		out[key] = append(slice, x)
	}
	return out
}
