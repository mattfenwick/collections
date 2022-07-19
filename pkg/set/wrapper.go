package set

import (
	"fmt"
	"github.com/mattfenwick/collections/pkg/function"
	"golang.org/x/exp/maps"
)

type Wrapper[A any, K comparable] struct {
	Elems      map[K]A
	Projection func(A) K
}

func NewWrapper[A comparable](elems []A) *Wrapper[A, A] {
	return NewWrapperBy(function.Id[A], elems)
}

func NewWrapperBy[A any, K comparable](projection func(A) K, elems []A) *Wrapper[A, K] {
	s := &Wrapper[A, K]{Elems: map[K]A{}, Projection: projection}
	for _, e := range elems {
		s.Add(e)
	}
	fmt.Printf("lengths? %d, %d\n", len(elems), len(s.Elems))
	return s
}

func (s *Wrapper[A, K]) Add(a A) bool {
	if !s.Contains(a) {
		s.Elems[s.Projection(a)] = a
		return true
	}
	return false
}

func (s *Wrapper[A, K]) Delete(a A) bool {
	_, ok := s.Elems[s.Projection(a)]
	delete(s.Elems, s.Projection(a))
	return ok
}

func (s *Wrapper[A, K]) Contains(a A) bool {
	_, ok := s.Elems[s.Projection(a)]
	return ok
}

func (s *Wrapper[A, K]) Len() int {
	return len(s.Elems)
}

func (s *Wrapper[A, K]) ToSlice() []A {
	return maps.Values(s.Elems)
}
