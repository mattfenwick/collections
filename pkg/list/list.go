package list

type List[A any] struct {
	Head A
	Tail *List[A] // TODO use maybe
}

func NewList[A any](xs []A) *List[A] {
	first := &List[A]{}
	last := first
	for _, x := range xs {
		next := &List[A]{Head: x}
		last.Tail = next
		last = next
	}
	return first.Tail
}
