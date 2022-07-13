package pkg

import (
	"fmt"
)

//func Equalable[T comparable](a T, b T) bool {
//	return a == b
//}
//
//func GreaterThan[T comparable](a T, b T) bool {
//	return a > b
//}

//type A interface {
//	~int | ~uint
//}
//
//type B[Q A] interface {
//	Q
//}
//
//type Comparable[T comparable] struct {
//	T T
//}
//
//func (c *Comparable[T]) Equal(other *Comparable[T]) bool {
//	return c.T == other.T
//}

func Example() {
	EqExample()
}

func EqExample() {
	//auints := []uint{1,2,3,4,5}
	a := []Uint{1, 2, 3, 4, 5}
	b := []Uint{0, 2, 4, 6, 8}
	for _, x := range b {
		fmt.Printf("looking for %d: result %d\n", x, Index(a, x))
	}

	fmt.Printf("Eq? %+v, %+v, %+v, %+v\n",
		SliceEq[Uint](a).Equal(a),
		SliceEq[Uint](a).Equal(b),
		SliceEq[Uint](b).Equal(a),
		SliceEq[Uint](b).Equal(b))

	ints := []Int{18, 27, 3, 39, -8, 37, 5, 12}
	//sorted := MergeSortWithComparator(ints, func(a int, b int) Ordering {
	//	if a < b {
	//		return OrderingLessThan
	//	} else if a == b {
	//		return OrderingEqual
	//	} else {
	//		return OrderingGreaterThan
	//	}
	//})
	sorted := Sort(ints)
	fmt.Printf("ints: %+v\nsorted: %+v\n", ints, sorted)
}
