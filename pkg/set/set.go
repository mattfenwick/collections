package set

import (
	"github.com/mattfenwick/collections/pkg/function"
	"github.com/pkg/errors"
)

type request[A any] struct {
	add      *A
	delete   *A
	contains *A
	len      bool
	toSlice  bool
}

type response[A any] struct {
	add      bool
	delete   bool
	contains bool
	len      int
	toSlice  []A
}

type Set[A any] struct {
	closure func(*request[A]) *response[A]
}

func NewSet[A comparable](elems []A) *Set[A] {
	return NewSetBy(function.Id[A], elems)
}

func NewSetBy[A any, K comparable](projection func(A) K, elems []A) *Set[A] {
	w := NewWrapperBy(projection, elems)
	s := &Set[A]{closure: func(r *request[A]) *response[A] {
		if r.add != nil {
			return &response[A]{add: w.Add(*r.add)}
		} else if r.delete != nil {
			return &response[A]{delete: w.Delete(*r.delete)}
		} else if r.contains != nil {
			return &response[A]{contains: w.Contains(*r.contains)}
		} else if r.len {
			return &response[A]{len: w.Len()}
		} else if r.toSlice {
			return &response[A]{toSlice: w.ToSlice()}
		}
		panic(errors.Errorf("invalid request: %+v", r))
	}}
	return s
}

func (s *Set[A]) Add(a A) bool {
	return s.closure(&request[A]{add: &a}).add
}

func (s *Set[A]) Delete(a A) bool {
	return s.closure(&request[A]{delete: &a}).delete
}

func (s *Set[A]) Contains(a A) bool {
	return s.closure(&request[A]{contains: &a}).contains
}

func (s *Set[A]) Len() int {
	return s.closure(&request[A]{len: true}).len
}

func (s *Set[A]) ToSlice() []A {
	return s.closure(&request[A]{toSlice: true}).toSlice
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
