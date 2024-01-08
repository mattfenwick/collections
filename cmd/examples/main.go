package main

import (
	"encoding/json"
	"fmt"
	. "github.com/mattfenwick/collections/pkg/base"
	"github.com/mattfenwick/collections/pkg/dict"
	"github.com/mattfenwick/collections/pkg/slice"
)

func main() {
	EqExample()
	SortExample()
}

func SortExample() {
	someInts := []int{4, 79, 13, -8, 22, 4, 8, 7}
	fmt.Printf("sort ints: %+v\n  %+v -- comparable\n  %+v -- Comparator\n  %+v -- Ord\n\n",
		someInts,
		slice.Sort(someInts),
		slice.SortOnBy(WrapInt, Compare[Int], someInts),
		slice.SortOn(WrapInt, someInts))

	fmt.Printf("sort a slice of slices: %+v\n\n",
		slice.SortBy(slice.CompareSlicePairwise[int](), [][]int{
			{3, 4, 5},
			{3, 4},
			{1, 2, 3},
			{},
		}))

	fmt.Printf("sort a slice of maps: %+v\n\n",
		slice.SortBy(dict.CompareMapPairwise[string, int](), []map[string]int{
			{"a": 1, "b": 2},
			{"a": 2, "b": 2},
			{"a": 2, "b": 1},
			{"a": 2, "b": 4},
			{"a": 1},
			{"a": 2, "b": 2, "c": 1},
			{"a": 1, "b": 2, "c": 3},
			{"a": 1, "b": 3},
			{},
		}))

	ints := []Int{18, 27, 3, 39, -8, 37, 5, 12}
	sorted := slice.Sort(ints)
	fmt.Printf("Ints: %+v\n  sorted: %+v\n\n", ints, sorted)

	pairs := []*Pair[int, string]{
		{Fst: 18, Snd: "jkl"},
		{Fst: 14, Snd: "ghi"},
		{Fst: 14, Snd: "tuv"},
		{Fst: 14, Snd: "jkl"},
		{Fst: 11, Snd: "abc"},
		{Fst: 17, Snd: "xyz"},
		{Fst: 13, Snd: "qrs"},
		{Fst: 16, Snd: "def"},
		{Fst: 10, Snd: "jkl"},
	}
	fmt.Printf("sort pairs: %+v\n  %+v -- natural\n  %+v -- first element\n  %+v -- 2nd element\n\n",
		DumpJson(pairs),
		DumpJson(slice.SortBy(ComparePair[int, string](), pairs)),
		DumpJson(slice.SortOn(Fst[int, string], pairs)),
		DumpJson(slice.SortOn(Snd[int, string], pairs)))
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
		fmt.Printf("looking for %d: result %d\n", x, slice.IndexEq(a, x))
	}

	fmt.Printf("Eq? %+v, %+v\n%+v, %+v, %+v, %+v\n",
		a, b,
		slice.SliceEq[Uint](a).Equal(a),
		slice.SliceEq[Uint](a).Equal(b),
		slice.SliceEq[Uint](b).Equal(a),
		slice.SliceEq[Uint](b).Equal(b))
	fmt.Println()
}
