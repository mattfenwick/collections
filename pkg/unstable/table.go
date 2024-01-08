package unstable

import (
	"github.com/mattfenwick/collections/pkg/base"
	"github.com/mattfenwick/collections/pkg/dict"
	"github.com/mattfenwick/collections/pkg/function"
	"github.com/mattfenwick/collections/pkg/iterable"
	"github.com/mattfenwick/collections/pkg/slice"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
)

type Table[A any, V any] struct {
	Set     func(A, V) bool
	Delete  func(A) *V
	Get     func(A) *V
	Len     func() int
	ToSlice func() []*base.Pair[A, V]
	Merge   func(*Table[A, V]) *Table[A, V]
	//Intersect  func(*Table[A, V]) *Table[A, *Pair[V, V]]
	//Difference func(*Table[A, V]) *Table[A, V]
	Iterator func() iterable.Iterator[*base.Pair[A, V]]
}

func Empty[A comparable, V any]() *Table[A, V] {
	return NewTableBy[A, A, V](function.Id[A], slice.Slice[*base.Pair[A, V]](nil))
}

func FromSlice[A comparable, V any](elems []*base.Pair[A, V]) *Table[A, V] {
	return NewTableBy[A, A, V](function.Id[A], slice.Slice[*base.Pair[A, V]](elems))
}

func FromSliceBy[A any, K comparable, V any](projection func(A) K, elems []*base.Pair[A, V]) *Table[A, V] {
	return NewTableBy[A, K, V](projection, slice.Slice[*base.Pair[A, V]](elems))
}

func NewTable[A comparable, V any](elems iterable.Iterable[*base.Pair[A, V]]) *Table[A, V] {
	return NewTableBy(function.Id[A], elems)
}

// NewTableBy allows creation of a table from an element type which isn't comparable.
//
//	Why is 'comparable' required?  Because this table implementation uses a map.
//	Unfortunately, it's not possible to add implementations of user-defined types for comparable.
//	This uses a projection function to create a `comparable` key.
//	VERY VERY VERY IMPORTANT NOTES about the projection function:
//	- it should be a pure function (no side effects)
//	- it should make intuitive sense
//	- don't mix tables which use different projections -- the results will be unpredictable and won't make sense
func NewTableBy[A any, K comparable, V any](projection func(A) K, initialElements iterable.Iterable[*base.Pair[A, V]]) *Table[A, V] {
	elems := map[K]*base.Pair[A, V]{}
	var t *Table[A, V]
	t = &Table[A, V]{
		// Set returns:
		//  - true: element was *not* already in table
		//  - false: element *was* already in table
		Set: func(a A, v V) bool {
			key := projection(a)
			_, ok := elems[key]
			elems[key] = base.NewPair(a, v)
			return !ok
		},
		Delete: func(a A) *V {
			pair, ok := elems[projection(a)]
			if ok {
				delete(elems, projection(a))
				return &pair.Snd
			}
			return nil
		},
		Get: func(a A) *V {
			pair, ok := elems[projection(a)]
			if !ok {
				return nil
			}
			return &pair.Snd
		},
		Len: func() int {
			return len(elems)
		},
		ToSlice: func() []*base.Pair[A, V] {
			return maps.Values(elems)
		},
		// Merge is right-biased, meaning: if a key is in both left and right, the value right
		//   value will be taken.  This means that in general, a.Merge(b) != b.Merge(a)
		Merge: func(other *Table[A, V]) *Table[A, V] {
			return FromSliceBy[A, K, V](projection, append(t.ToSlice(), other.ToSlice()...))
		},
		Iterator: func() iterable.Iterator[*base.Pair[A, V]] {
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
			t.Set((*x).Fst, (*x).Snd)
		}
	}
	return t
}

func (t *Table[A, V]) Contains(a A) bool {
	return t.Get(a) != nil
}

// Entries returns a sorted slice of entries, using the natural sort order
//
//	of the table's keys
func Entries[A constraints.Ordered, V any](t *Table[A, V]) []*base.Pair[A, V] {
	return slice.SortOn(base.Fst[A, V], t.ToSlice())
}
