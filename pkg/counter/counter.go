package counter

import (
	"github.com/mattfenwick/collections/pkg/base"
	"github.com/mattfenwick/collections/pkg/builtin"
	"github.com/mattfenwick/collections/pkg/dict"
	"github.com/mattfenwick/collections/pkg/function"
	"github.com/mattfenwick/collections/pkg/iterable"
	"github.com/mattfenwick/collections/pkg/slice"
	"golang.org/x/exp/maps"
)

type Counter[A any] struct {
	AddN           func(A, int) int
	RemoveN        func(A, int) int
	Get            func(A) int
	Len            func() int
	ToSlice        func() []*base.Pair[A, int]
	Union          func(*Counter[A]) *Counter[A]
	Intersect      func(*Counter[A]) *Counter[A]
	Difference     func(*Counter[A]) *Counter[A]
	Iterator       func() iterable.Iterator[*base.Pair[A, int]]
	KeysIterator   func() iterable.Iterator[A]
	CountsIterator func() iterable.Iterator[int]
}

func FromSlice[A comparable](elems []*base.Pair[A, int]) *Counter[A] {
	return NewCounterBy[A](function.Id[A], slice.Slice[*base.Pair[A, int]](elems))
}

func FromSliceBy[A any, K comparable](projection func(A) K, elems []*base.Pair[A, int]) *Counter[A] {
	return NewCounterBy[A, K](projection, slice.Slice[*base.Pair[A, int]](elems))
}

func NewCounter[A comparable](elems iterable.Iterable[*base.Pair[A, int]]) *Counter[A] {
	return NewCounterBy(function.Id[A], elems)
}

func NewCounterBy[A any, K comparable](projection func(A) K, initialElements iterable.Iterable[*base.Pair[A, int]]) *Counter[A] {
	counts := map[K]*base.Pair[A, int]{}
	var c *Counter[A]
	c = &Counter[A]{
		AddN: func(a A, increment int) int {
			key := projection(a)
			pair, ok := counts[key]
			newValue := increment
			if ok {
				newValue += pair.Snd
			}
			counts[key] = base.NewPair(pair.Fst, newValue)
			return newValue
		},
		RemoveN: func(a A, increment int) int {
			key := projection(a)
			pair, ok := counts[key]
			if !ok {
				return 0
			}
			newValue := pair.Snd - increment
			if newValue <= 0 {
				delete(counts, key)
				return 0
			}
			counts[key] = base.NewPair(pair.Fst, newValue)
			return newValue
		},
		Get: func(a A) int {
			pair, ok := counts[projection(a)]
			if !ok {
				return 0
			}
			return pair.Snd
		},
		Len: func() int { return len(counts) },
		ToSlice: func() []*base.Pair[A, int] {
			return maps.Values(counts)
		},
		Union: func(other *Counter[A]) *Counter[A] {
			return FromSliceBy[A, K](projection, append(c.ToSlice(), other.ToSlice()...))
		},
		Intersect: func(other *Counter[A]) *Counter[A] {
			var out []*base.Pair[A, int]
			for _, pair := range counts {
				otherVal := other.Get(pair.Fst)
				newVal := builtin.Min(pair.Snd, otherVal)
				if newVal > 0 {
					out = append(out, base.NewPair(pair.Fst, newVal))
				}
			}
			return FromSliceBy[A, K](projection, out)
		},
		Difference: func(other *Counter[A]) *Counter[A] {
			var out []*base.Pair[A, int]
			for _, pair := range counts {
				otherVal := other.Get(pair.Fst)
				newVal := pair.Snd - otherVal
				if newVal > 0 {
					out = append(out, base.NewPair(pair.Fst, newVal))
				}
			}
			return FromSliceBy[A, K](projection, out)
		},
		Iterator: func() iterable.Iterator[*base.Pair[A, int]] {
			return dict.ValuesIterator(counts)
		},
		KeysIterator: func() iterable.Iterator[A] {
			return iterable.Map(base.Fst[A, int], c.Iterator())
		},
		CountsIterator: func() iterable.Iterator[int] {
			return iterable.Map(base.Snd[A, int], c.Iterator())
		},
	}
	iterator := initialElements.Iterator()
	for {
		x := iterator.Next()
		if x == nil {
			break
		}
		c.AddN((*x).Fst, (*x).Snd)
	}
	return c
}

func (c *Counter[A]) Add(a A) int {
	return c.AddN(a, 1)
}

func (c *Counter[A]) Remove(a A) int {
	return c.RemoveN(a, 1)
}
