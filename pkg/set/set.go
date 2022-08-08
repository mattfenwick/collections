package set

import (
	"github.com/mattfenwick/collections/pkg/dict"
	"github.com/mattfenwick/collections/pkg/function"
	"github.com/mattfenwick/collections/pkg/iterable"
	"github.com/mattfenwick/collections/pkg/slice"
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
	Iterator   func() iterable.Iterator[A]
}

func FromSlice[A comparable](elems []A) *Set[A] {
	return NewSetBy[A](function.Id[A], slice.Slice[A](elems))
}

func FromSliceBy[A any, K comparable](projection func(A) K, elems []A) *Set[A] {
	return NewSetBy[A, K](projection, slice.Slice[A](elems))
}

func NewSet[A comparable](elems iterable.Iterable[A]) *Set[A] {
	return NewSetBy(function.Id[A], elems)
}

func NewSetBy[A any, K comparable](projection func(A) K, initialElements iterable.Iterable[A]) *Set[A] {
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
			return NewSetBy[A, K](projection, slice.Slice[A](append(s.ToSlice(), other.ToSlice()...)))
		},
		Intersect: func(other *Set[A]) *Set[A] {
			var out []A
			for _, val := range elems {
				if other.Contains(val) {
					out = append(out, val)
				}
			}
			return NewSetBy[A, K](projection, slice.Slice[A](out))
		},
		Difference: func(other *Set[A]) *Set[A] {
			var out []A
			for _, val := range elems {
				if !other.Contains(val) {
					out = append(out, val)
				}
			}
			return NewSetBy[A, K](projection, slice.Slice[A](out))
		},
		Iterator: func() iterable.Iterator[A] {
			return dict.ValuesIterator(elems)
		},
	}
	iterator := initialElements.Iterator()
	for {
		x := iterator.Next()
		if x == nil {
			break
		}
		s.Add(*x)
	}
	return s
}
