package dict

import (
	. "github.com/mattfenwick/collections/pkg/base"
	"github.com/mattfenwick/collections/pkg/slice"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
)

func ToSlice[A comparable, B any](m map[A]B) []*Pair[A, B] {
	return slice.Map(func(key A) *Pair[A, B] {
		return NewPair(key, m[key])
	}, maps.Keys(m))
}

func FromSliceBy[A comparable, B any](merge func(B, B) B, ps []*Pair[A, B]) map[A]B {
	out := map[A]B{}
	for _, p := range ps {
		if val, ok := out[p.Fst]; ok {
			out[p.Fst] = merge(val, p.Snd)
		} else {
			out[p.Fst] = p.Snd
		}
	}
	return out
}

func Map[A comparable, B, C any](f func(B) C, xs map[A]B) map[A]C {
	return MapWithKey(func(a A, b B) C { return f(b) }, xs)
}

func MapWithKey[A comparable, B, C any](f func(A, B) C, xs map[A]B) map[A]C {
	out := map[A]C{}
	for k, v := range xs {
		out[k] = f(k, v)
	}
	return out
}

func Merge[A comparable, B constraints.Ordered](m1 map[A]B, m2 map[A]B) map[A]B {
	return MergeBy(func(a A, b1 B, b2 B) B {
		if b1 >= b2 {
			return b1
		}
		return b2
	}, m1, m2)
}

func MergeOrd[A comparable, B Ord[B]](m1 map[A]B, m2 map[A]B) map[A]B {
	return MergeBy(func(a A, b1 B, b2 B) B {
		if GreaterThanOrEqual(b1, b2) {
			return b1
		}
		return b2
	}, m1, m2)
}

func MergeBy[A comparable, B any](resolve func(A, B, B) B, m1 map[A]B, m2 map[A]B) map[A]B {
	out := map[A]B{}
	for k, v1 := range m1 {
		if v2, ok := m2[k]; ok {
			// if k is in *both* m1 and m2, use the resolve function to decide the value
			out[k] = resolve(k, v1, v2)
		} else {
			out[k] = v1
		}
	}
	// from m2: only add kvs that *aren't* in m1
	for k, v2 := range m2 {
		if _, ok := m1[k]; !ok {
			out[k] = v2
		}
	}
	return out
}
