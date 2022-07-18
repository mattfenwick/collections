package pkg

import (
	"github.com/mattfenwick/collections/pkg/functions"
	"golang.org/x/exp/maps"
)

type Set[A any, K comparable] struct {
	Elems      map[K]A
	Projection func(A) K
}

func NewSet[A comparable](elems []A) *Set[A, A] {
	return NewSetBy(functions.Id[A], elems)
}

func NewSetBy[A any, K comparable](projection func(A) K, elems []A) *Set[A, K] {
	s := &Set[A, K]{Elems: map[K]A{}, Projection: projection}
	for _, e := range elems {
		s.Add(e)
	}
	return s
}

func (s *Set[A, K]) Add(a A) {
	s.Elems[s.Projection(a)] = a
}

func (s *Set[A, K]) Delete(a A) {
	delete(s.Elems, s.Projection(a))
}

func (s *Set[A, K]) Contains(a A) bool {
	_, ok := s.Elems[s.Projection(a)]
	return ok
}

func (s *Set[A, K]) Len() int {
	return len(s.Elems)
}

func (s *Set[A, K]) ToSlice() []A {
	return maps.Values(s.Elems)
}
