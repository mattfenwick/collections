package pkg

// Unit represents the Haskell value `()`
type Unit struct{}

var UnitC = &Unit{}

type Pair[A, B any] struct {
	A A
	B B
}

func NewPair[A, B any](a A, b B) *Pair[A, B] {
	return &Pair[A, B]{A: a, B: b}
}

func First[A, B any](p *Pair[A, B]) A {
	return p.A
}

func Second[A, B any](p *Pair[A, B]) B {
	return p.B
}
