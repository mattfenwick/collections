package dict

import (
	. "github.com/mattfenwick/collections/pkg/base"
	"github.com/mattfenwick/collections/pkg/builtin"
	"golang.org/x/exp/maps"
)

func EqualIndexEq[A comparable, B Eq[B]](key A) Equaler[map[A]B] {
	return EqualIndexBy(key, Equal[B])
}

func EqualIndex[A comparable, B comparable](key A) Equaler[map[A]B] {
	return EqualIndexBy(key, builtin.EQ[B])
}

// EqualIndexBy equals a single index
func EqualIndexBy[A comparable, B any](key A, equal Equaler[B]) Equaler[map[A]B] {
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

func EqualPairwiseEq[A comparable, B Eq[B]]() Equaler[map[A]B] {
	return EqualPairwiseBy[A, B](Equal[B])
}

func EqualPairwise[A comparable, B comparable]() Equaler[map[A]B] {
	return EqualPairwiseBy[A, B](builtin.EQ[B])
}

// EqualPairwiseBy works by project a map to a list, therefore it's inefficient and probably
//
//	best to avoid using unless absolutely necessary!
func EqualPairwiseBy[A comparable, B any](equal Equaler[B]) Equaler[map[A]B] {
	return func(xs map[A]B, ys map[A]B) bool {
		return maps.EqualFunc(xs, ys, equal)
	}
}
