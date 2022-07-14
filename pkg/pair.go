package pkg

type Pair[A, B any] struct {
	Fst A
	Snd B
}

func NewPair[A, B any](a A, b B) *Pair[A, B] {
	return &Pair[A, B]{Fst: a, Snd: b}
}

func First[A, B any](p *Pair[A, B]) A {
	return p.Fst
}

func Second[A, B any](p *Pair[A, B]) B {
	return p.Snd
}

func EqualPair[A Eq[A], B Eq[B]](p1 *Pair[A, B], p2 *Pair[A, B]) bool {
	return p1.Fst.Equal(p2.Fst) && p1.Snd.Equal(p2.Snd)
}

// TODO is there any way to equip Pair with an Eq instance?
//func (p *Pair[A, B]) Equal(p2 *Pair[A, B]) bool {
//	return p.Fst.Equal(p2.Fst)
//}

type PairEq[A Eq[A], B Eq[B]] Pair[A, B]

func (p1 PairEq[A, B]) Equal(p2 PairEq[A, B]) bool {
	return p1.Fst.Equal(p2.Fst) && p1.Snd.Equal(p2.Snd)
}

type PairOrd[A Ord[A], B Ord[B]] Pair[A, B]

func (p1 PairOrd[A, B]) Equal(p2 PairOrd[A, B]) bool {
	return PairEq[A, B](p1).Equal(PairEq[A, B](p2))
}

func (p1 PairOrd[A, B]) Compare(p2 PairOrd[A, B]) Ordering {
	if comp := p1.Fst.Compare(p2.Fst); comp != OrderingEqual {
		return comp
	}
	return p1.Snd.Compare(p2.Snd)
}
