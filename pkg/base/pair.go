package base

import (
	"golang.org/x/exp/constraints"
)

type Pair[A, B any] struct {
	Fst A
	Snd B
}

func NewPair[A, B any](a A, b B) *Pair[A, B] {
	return &Pair[A, B]{Fst: a, Snd: b}
}

func Fst[A, B any](p *Pair[A, B]) A {
	return p.Fst
}

func Snd[A, B any](p *Pair[A, B]) B {
	return p.Snd
}

// TODO is there any way to equip Pair with an Eq instance?
//func (p *Pair[A, B]) Equal(p2 *Pair[A, B]) bool {
//	return p.Fst.Equal(p2.Fst)
//}

func EqualPairEq[A Eq[A], B Eq[B]]() Equaler[*Pair[A, B]] {
	return EqualPairBy(Equal[A], Equal[B])
}

func EqualPair[A comparable, B comparable]() Equaler[*Pair[A, B]] {
	return EqualPairBy(EqualComparable[A], EqualComparable[B])
}

func EqualPairBy[A, B any](fst Equaler[A], snd Equaler[B]) Equaler[*Pair[A, B]] {
	return func(p1 *Pair[A, B], p2 *Pair[A, B]) bool {
		return fst(p1.Fst, p2.Fst) && snd(p1.Snd, p2.Snd)
	}
}

func ComparePairOrd[A Ord[A], B Ord[B]]() Comparator[*Pair[A, B]] {
	return ComparePairBy(Compare[A], Compare[B])
}

func ComparePair[A constraints.Ordered, B constraints.Ordered]() Comparator[*Pair[A, B]] {
	return ComparePairBy(CompareOrdered[A], CompareOrdered[B])
}

func ComparePairBy[A, B any](fst Comparator[A], snd Comparator[B]) Comparator[*Pair[A, B]] {
	return func(p1 *Pair[A, B], p2 *Pair[A, B]) Ordering {
		c := fst(p1.Fst, p2.Fst)
		if c != OrderingEqual {
			return c
		}
		return snd(p1.Snd, p2.Snd)
	}
}
