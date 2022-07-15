package pkg

import "golang.org/x/exp/maps"
import . "github.com/mattfenwick/collections/pkg/base"

type Set[T EqOrComparable[T]] struct {
	Elems map[T]bool
}

func NewSet[T EqOrComparable[T]](elems []T) *Set[T] {
	s := &Set[T]{Elems: map[T]bool{}}
	for _, e := range elems {
		s.Add(e)
	}
	return s
}

func (s *Set[T]) Add(a T) {
	s.Elems[a] = true
}

func (s *Set[T]) Delete(a T) {
	delete(s.Elems, a)
}

func (s *Set[T]) Contains(a T) bool {
	return s.Elems[a]
}

func (s *Set[T]) Len() int {
	return len(s.Elems)
}

func (s *Set[T]) ToSlice() []T {
	return maps.Keys(s.Elems)
}
