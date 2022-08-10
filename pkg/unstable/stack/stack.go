package stack

import (
	"github.com/mattfenwick/collections/pkg/iterable"
	"github.com/mattfenwick/collections/pkg/slice"
)

type Stack[A any] struct {
	Items []A
}

func Empty[A any]() *Stack[A] {
	return FromSlice[A](nil)
}

func FromSlice[A any](initialElements []A) *Stack[A] {
	return FromIterator[A](slice.Slice[A](initialElements))
}

func FromIterator[A any](initialElements iterable.Iterable[A]) *Stack[A] {
	s := &Stack[A]{}
	iterator := initialElements.Iterator()
	for {
		x := iterator.Next()
		if x == nil {
			break
		}
		s.Push(*x)
	}
	return s
}

func (s *Stack[A]) Peek() *A {
	if len(s.Items) == 0 {
		return nil
	}
	return &s.Items[len(s.Items)-1]
}

func (s *Stack[A]) Push(a A) {
	s.Items = append(s.Items, a)
}

func (s *Stack[A]) Pop() *A {
	if len(s.Items) == 0 {
		return nil
	}
	next := s.Items[len(s.Items)-1]
	s.Items = s.Items[:len(s.Items)-1]
	return &next
}

func (s *Stack[A]) Iterator() iterable.Iterator[A] {
	// TODO this is backwards from initialization -- is that a bad thing?
	return slice.Iterator(s.Items)
}
