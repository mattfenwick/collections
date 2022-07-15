package pkg

import . "github.com/mattfenwick/collections/pkg/base"

// for golang examples, see:
//   https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#map_reduce_filter for inspiration

// these functions are inspired by Haskell's Data.List:
//   https://hackage.haskell.org/package/base-4.16.2.0/docs/Data-List.html

func FilterSlice[A any](f F1[A, bool], xs []A) []A {
	var out []A
	for _, x := range xs {
		if f(x) {
			out = append(out, x)
		}
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

func GroupSlice[A EqOrComparable[A]](xs []A) map[A][]A {
	out := map[A][]A{}
	for _, x := range xs {
		slice := out[x]
		out[x] = append(slice, x)
	}
	return out
}

func PartitionSlice[A any](predicate F1[A, bool], xs []A) ([]A, []A) {
	var yes, no []A
	for _, x := range xs {
		if predicate(x) {
			yes = append(yes, x)
		} else {
			no = append(no, x)
		}
	}
	return yes, no
}
