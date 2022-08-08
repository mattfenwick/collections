package set

import (
	"github.com/mattfenwick/collections/pkg/function"
	"github.com/pkg/errors"
	"golang.org/x/exp/maps"
)

type Set[A any] struct {
	add      func(A) bool
	delete   func(A) bool
	contains func(A) bool
	length   func() int
	toSlice  func() []A
}

func NewSet[A comparable](elems []A) *Set[A] {
	return NewSetBy(function.Id[A], elems)
}

func NewSetBy[A any, K comparable](projection func(A) K, initialElements []A) *Set[A] {
	elems := map[K]A{}
	var s *Set[A]
	s = &Set[A]{
		add: func(a A) bool {
			if !s.Contains(a) {
				elems[projection(a)] = a
				return true
			}
			return false
		},
		delete: func(a A) bool {
			_, ok := elems[projection(a)]
			delete(elems, projection(a))
			return ok
		},
		contains: func(a A) bool {
			_, ok := elems[projection(a)]
			return ok
		},
		length: func() int {
			return len(elems)
		},
		toSlice: func() []A {
			return maps.Values(elems)
		},
	}
	for _, x := range initialElements {
		s.Add(x)
	}
	return s
}

func (s *Set[A]) Add(a A) bool {
	return s.add(a)
}

func (s *Set[A]) Delete(a A) bool {
	return s.delete(a)
}

func (s *Set[A]) Contains(a A) bool {
	return s.contains(a)
}

func (s *Set[A]) Len() int {
	return s.length()
}

func (s *Set[A]) ToSlice() []A {
	return s.toSlice()
}

func (s *Set[A]) Union(other *Set[A]) int {
	added := 0
	for _, e := range other.ToSlice() {
		if s.Add(e) {
			added++
		}
	}
	return added
}

// Intersect removes any items from `s` which are NOT in `other`
func (s *Set[A]) Intersect(other *Set[A]) int {
	removed := 0
	for _, e := range s.ToSlice() {
		if !other.Contains(e) {
			if ok := s.Delete(e); !ok {
				panic(errors.Errorf("unable to delete element %+v", e))
			}
			removed++
		}
	}
	return removed
}

// Difference removes any items from `s` which are in `other`
func (s *Set[A]) Difference(other *Set[A]) int {
	removed := 0
	for _, e := range s.ToSlice() {
		if other.Contains(e) {
			if ok := s.Delete(e); !ok {
				panic(errors.Errorf("unable to delete element %+v from set", e))
			}
			removed++
		}
	}
	return removed
}
