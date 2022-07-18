package maps

import (
	. "github.com/mattfenwick/collections/pkg/base"
	"github.com/mattfenwick/collections/pkg/builtins"
	"golang.org/x/exp/maps"
)

func EqualMapIndexEq[A comparable, B Eq[B]](key A) Equaler[map[A]B] {
	return EqualMapIndexBy(key, Equal[B])
}

func EqualMapIndex[A comparable, B comparable](key A) Equaler[map[A]B] {
	return EqualMapIndexBy(key, builtins.Equal[B])
}

// EqualMapIndexBy equals a single index
func EqualMapIndexBy[A comparable, B any](key A, equal Equaler[B]) Equaler[map[A]B] {
	return func(xs map[A]B, ys map[A]B) bool {
		x, xok := xs[key]
		y, yok := ys[key]
		if xok && yok {
			// both have the key?  check the value's equality
			return equal(x, y)
		} else if !xok && !yok {
			// neither has it
			return true
		}
		// one has it and the other doesn't
		return false
	}
}

func EqualMapPairwiseEq[A comparable, B Eq[B]]() Equaler[map[A]B] {
	return EqualMapPairwiseBy[A, B](Equal[B])
}

func EqualMapPairwise[A comparable, B comparable]() Equaler[map[A]B] {
	return EqualMapPairwiseBy[A, B](builtins.Equal[B])
}

// EqualMapPairwiseBy works by project a map to a list, therefore it's inefficient and probably
//   best to avoid using unless absolutely necessary!
func EqualMapPairwiseBy[A comparable, B any](equal Equaler[B]) Equaler[map[A]B] {
	return func(xs map[A]B, ys map[A]B) bool {
		return maps.EqualFunc(xs, ys, equal)
	}
}
