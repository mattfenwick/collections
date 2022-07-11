package doesntwork

/*
import (
	"fmt"
)

type Ordered interface {
	~int | ~uint | ~int8 | ~uint8 | ~int16 | ~uint16 |
		~int32 | ~uint32 | ~int64 | ~uint64 | ~uintptr |
		~float32 | ~float64 | ~string
}

type OrderedWrapper[T Ordered] struct {
	T T
}

func Wrap[T Ordered](t T) *OrderedWrapper[T] {
	return &OrderedWrapper[T]{t}
}

func (a *OrderedWrapper[T]) Equal(b *OrderedWrapper[T]) bool {
	return a.T == b.T
}

func (a *OrderedWrapper[T]) Compare(b *OrderedWrapper[T]) Ordering {
	if a.T < b.T {
		return OrderingLessThan
	} else if a.T == b.T {
		return OrderingEqual
	} else {
		return OrderingGreaterThan
	}
}

func Eq2[A Ordered](a A, b A) bool {
	return a == b
}

func Comp2[A Ordered](a A, b A) Ordering {
	if a < b {
		return OrderingLessThan
	} else if a == b {
		return OrderingEqual
	} else {
		return OrderingGreaterThan
	}
}

type Ordered2[T any] interface {
	Ordered
	Ord[T]
	Eq[T]
}

func Max2[S ~[]E, E Ord[E]](vs S) E {
	if len(vs) == 0 {
		panic("no elements")
	}

	var r = vs[0]
	for i := range vs[1:] {
		if GreaterThan(vs[i], r) {
			//if vs[i] > r {
			r = vs[i]
		}
	}
	return r
}

func Example() {
	unwrapped := []uint64{45, 27, 79, 3, 103, 4, 5, 6, 10, 9, 8}
	wrapped := MapSlice(unwrapped, Wrap[uint64])
	maxI := Max2(wrapped)
	fmt.Printf("max2: %+v , %+v\n%+v\n", unwrapped, wrapped, maxI)
}
*/
