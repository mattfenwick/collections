package queue

import (
	"github.com/mattfenwick/collections/pkg/iterable"
	"github.com/mattfenwick/collections/pkg/slice"
)

type Queue[A any] struct {
	Items []A
}

func Empty[A any]() *Queue[A] {
	return FromSlice[A](nil)
}

func FromSlice[A any](initialElements []A) *Queue[A] {
	return FromIterator[A](slice.Slice[A](initialElements))
}

func FromIterator[A any](initialElements iterable.Iterable[A]) *Queue[A] {
	s := &Queue[A]{}
	iterator := initialElements.Iterator()
	for {
		x := iterator.Next()
		if x == nil {
			break
		}
		s.PushBack(*x)
	}
	return s
}

func (s *Queue[A]) PeekFront() *A {
	if len(s.Items) == 0 {
		return nil
	}
	return &s.Items[0]
}

func (s *Queue[A]) PushBack(a A) {
	s.Items = append(s.Items, a)
}

func (s *Queue[A]) PopFront() *A {
	if len(s.Items) == 0 {
		return nil
	}
	next := s.Items[0]
	s.Items = s.Items[1:]
	return &next
}

func (s *Queue[A]) Iterator() iterable.Iterator[A] {
	return slice.Iterator(s.Items)
}
