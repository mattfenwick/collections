package main

import (
	"fmt"
	"github.com/mattfenwick/collections/pkg/base"
	"github.com/mattfenwick/collections/pkg/dict"
	"github.com/mattfenwick/collections/pkg/iterable"
	"github.com/mattfenwick/collections/pkg/json"
	"github.com/mattfenwick/collections/pkg/set"
	"github.com/mattfenwick/collections/pkg/slice"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Infof("starting")
	s1 := set.FromSlice([]int{1, 5, 2, 2, 7, 1})
	s2 := set.FromSlice([]int{4, 2, 7, 8, 2, 3})
	logrus.Infof("printing")
	fmt.Printf("%+v\n%+v\n%+v\n%+v\n", s1.Len(), s1.Union(s2), s1.Intersect(s2), s1.Difference(s2))

	underlying := map[string]int{"abc": 123, "def": 456, "ccc": 100}
	//d1 := dict.Dict[string, int](underlying)
	logrus.Infof("starting loop")
	for i := dict.KeysIterator(underlying); ; {
		v := i.Next()
		logrus.Infof("loop: %+v", json.MustMarshalToString(v))
		if v == nil {
			break
		}
	}
	logrus.Infof("finished loop")

	xss := [][]string{
		nil,
		{},
		{"abc"},
	}
	json.Print(iterable.ToSlice(slice.Iterator(xss)))

	yss := []*base.Pair[int, int]{
		nil,
		base.NewPair(1, 2),
		nil,
		nil,
	}
	json.Print(iterable.ToSlice(slice.Iterator(yss)))
}

type MyEqualer[A any] interface {
	Equaler() base.Equaler[A]
}

type P[A, B any] struct {
	Fst A
	Snd B
}

func Fst[A, B any](p *P[A, B]) A {
	return p.Fst
}

func Snd[A, B any](p *P[A, B]) B {
	return p.Snd
}

//func (p *P[A, B]) Equaler() base.Equaler[*P[A, B]] {
//	return slices.EqualBy[*P[A, B]](
//		functions.On(base.Equal[A], Fst[A, B]),
//		functions.On(base.Equal[B], Snd[A, B]))
//}

//func R() {
//	var q base.Equaler[*P[int, bool]]
//	p := &P[int, bool]{3, true}
//	q = p
//	Q[*P[int, bool]](p)
//}
//
//func Q[A base.Equaler[A]](x A) {
//
//}
