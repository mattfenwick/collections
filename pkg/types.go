package pkg

// Unit represents the Haskell value `()`
type Unit struct{}

var UnitC = &Unit{}

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
