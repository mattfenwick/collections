package main

import (
	"encoding/json"
	"fmt"
	. "github.com/mattfenwick/collections/pkg/base"
	"github.com/mattfenwick/collections/pkg/builtins"
	"github.com/mattfenwick/collections/pkg/slices"
)

func main() {
	EqExample()
	SortExample()
}

func SortExample() {
	someInts := []int{4, 79, 13, -8, 22, 4, 8, 7}
	fmt.Printf("sort ints: %+v\n  %+v -- comparable\n  %+v -- Comparator\n  %+v -- Ord\n\n",
		someInts,
		slices.SortBy(builtins.CompareOrdered[int], someInts),
		slices.SortOnBy(WrapInt, Compare[Int], someInts),
		slices.SortOn(WrapInt, someInts))

	fmt.Printf("sort a slice of slices: %+v\n\n",
		slices.SortBy(slices.CompareSlice(builtins.CompareOrdered[int]), [][]int{
			{3, 4, 5},
			{3, 4},
			{1, 2, 3},
			{},
		}))

	ints := []Int{18, 27, 3, 39, -8, 37, 5, 12}
	sorted := slices.Sort(ints)
	fmt.Printf("Ints: %+v\n  sorted: %+v\n\n", ints, sorted)

	pairs := []*Pair[int, string]{
		{18, "jkl"},
		{14, "ghi"},
		{14, "tuv"},
		{14, "jkl"},
		{11, "abc"},
		{17, "xyz"},
		{13, "qrs"},
		{16, "def"},
		{10, "jkl"},
	}
	fmt.Printf("sort pairs: %+v\n  %+v -- natural\n  %+v -- first element\n  %+v -- 2nd element\n\n",
		DumpJson(pairs),
		DumpJson(slices.SortBy(slices.ComparePairOrdered[int, string](), pairs)),
		DumpJson(slices.SortOnBy(First[int, string], builtins.CompareOrdered[int], pairs)),
		DumpJson(slices.SortOnBy(Second[int, string], builtins.CompareOrdered[string], pairs)))
}

func DumpJson(obj interface{}) string {
	bytes, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func EqExample() {
	a := []Uint{1, 2, 3, 4, 5}
	b := []Uint{0, 2, 4, 6, 8}
	for _, x := range b {
		fmt.Printf("looking for %d: result %d\n", x, Index(a, x))
	}

	fmt.Printf("Eq? %+v, %+v\n%+v, %+v, %+v, %+v\n",
		a, b,
		slices.SliceEq[Uint](a).Equal(a),
		slices.SliceEq[Uint](a).Equal(b),
		slices.SliceEq[Uint](b).Equal(a),
		slices.SliceEq[Uint](b).Equal(b))
	fmt.Println()
}
