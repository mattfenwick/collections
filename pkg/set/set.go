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

func Empty[A comparable]() *Set[A] {
	return NewSetBy[A](function.Id[A], slice.Slice[A](nil))
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

// NewSetBy allows creation of a set from an element type which isn't comparable.
//
//	Why is 'comparable' required?  Because this set implementation uses a map.
//	Unfortunately, it's not possible to add implementations of user-defined types for comparable.
//	This uses a projection function to create a `comparable` key.
//	VERY VERY VERY IMPORTANT NOTES about the projection function:
//	- it should be a pure function (no side effects)
//	- it should make intuitive sense
//	- don't mix sets which use different projections -- the results will be unpredictable and won't make sense
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
			return FromSliceBy[A, K](projection, append(s.ToSlice(), other.ToSlice()...))
		},
		Intersect: func(other *Set[A]) *Set[A] {
			var out []A
			for _, val := range elems {
				if other.Contains(val) {
					out = append(out, val)
				}
			}
			return FromSliceBy[A, K](projection, out)
		},
		Difference: func(other *Set[A]) *Set[A] {
			var out []A
			for _, val := range elems {
				if !other.Contains(val) {
					out = append(out, val)
				}
			}
			return FromSliceBy[A, K](projection, out)
		},
		Iterator: func() iterable.Iterator[A] {
			return dict.ValuesIterator(elems)
		},
	}
	if initialElements != nil {
		iterator := initialElements.Iterator()
		for {
			x := iterator.Next()
			if x == nil {
				break
			}
			s.Add(*x)
		}
	}
	return s
}
