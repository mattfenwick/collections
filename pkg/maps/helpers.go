package maps

import (
	. "github.com/mattfenwick/collections/pkg/base"
	"github.com/mattfenwick/collections/pkg/slices"
	"golang.org/x/exp/maps"
)

func ToSlice[A comparable, B any](m map[A]B) []*Pair[A, B] {
	return slices.Map(func(key A) *Pair[A, B] {
		return NewPair(key, m[key])
	}, maps.Keys(m))
}

func FromSlice[A comparable, B any](ps []*Pair[A, B]) map[A]B {
	panic("TODO")
}

func Merge[A comparable, B Ord[B]](m1 map[A]B, m2 map[A]B) map[A]B {
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
