package base

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
