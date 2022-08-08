package set

import (
	"github.com/mattfenwick/collections/pkg/function"
	"golang.org/x/exp/maps"
)

type Set[A any] struct {
	Add        func(A) bool
	Delete     func(A) bool
	Contains   func(A) bool
	Len        func() int
	ToSlice    func() []A
	Union      func(*Set[A]) *Set[A]
	Intersect  func(*Set[A]) *Set[A]
	Difference func(*Set[A]) *Set[A]
}

func NewSet[A comparable](elems []A) *Set[A] {
	return NewSetBy(function.Id[A], elems)
}

func NewSetBy[A any, K comparable](projection func(A) K, initialElements []A) *Set[A] {
	elems := map[K]A{}
	var s *Set[A]
	s = &Set[A]{
		Add: func(a A) bool {
			if !s.Contains(a) {
				elems[projection(a)] = a
				return true
			}
			return false
		},
		Delete: func(a A) bool {
			_, ok := elems[projection(a)]
			delete(elems, projection(a))
			return ok
		},
		Contains: func(a A) bool {
			_, ok := elems[projection(a)]
			return ok
		},
		Len: func() int {
			return len(elems)
		},
		ToSlice: func() []A {
			return maps.Values(elems)
		},
		Union: func(other *Set[A]) *Set[A] {
			return NewSetBy(projection, append(s.ToSlice(), other.ToSlice()...))
		},
		Intersect: func(other *Set[A]) *Set[A] {
			var out []A
			for _, val := range elems {
				if other.Contains(val) {
					out = append(out, val)
				}
			}
			return NewSetBy(projection, out)
		},
		Difference: func(other *Set[A]) *Set[A] {
			var out []A
			for _, val := range elems {
				if !other.Contains(val) {
					out = append(out, val)
				}
			}
			return NewSetBy(projection, out)
		},
	}
	for _, x := range initialElements {
		s.Add(x)
	}
	return s
}
