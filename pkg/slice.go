package pkg

import . "github.com/mattfenwick/collections/pkg/base"

// for golang examples, see:
//   https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#map_reduce_filter for inspiration

// these functions are inspired by Haskell's Data.List:
//   https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html

func GroupSlice[A EqOrComparable[A]](xs []A) map[A][]A {
	out := map[A][]A{}
	for _, x := range xs {
		slice := out[x]
		out[x] = append(slice, x)
	}
	return out
}
