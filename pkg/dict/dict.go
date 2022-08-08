package dict

import (
	"github.com/mattfenwick/collections/pkg/base"
	"github.com/mattfenwick/collections/pkg/iterable"
)

type Dict[K comparable, V any] map[K]V

func (d Dict[K, V]) Iterator() iterable.Iterator[*base.Pair[K, V]] {
	return Iterator(d)
}

func Iterator[K comparable, V any](xs map[K]V) iterable.Iterator[*base.Pair[K, V]] {
	c := make(chan *base.Pair[K, V])
	go func() {
		for k, v := range xs {
			c <- base.NewPair(k, v)
		}
		close(c)
	}()
	return &iterable.FunctionIterator[*base.Pair[K, V]]{
		F: func() **base.Pair[K, V] {
			val, ok := <-c
			if ok {
				return &val
			}
			return nil
		},
	}
}

func ValuesIterator[K comparable, V any](xs map[K]V) iterable.Iterator[V] {
	c := make(chan V)
	go func() {
		for _, v := range xs {
			c <- v
		}
		close(c)
	}()
	return &iterable.FunctionIterator[V]{
		F: func() *V {
			val, ok := <-c
			if ok {
				return &val
			}
			return nil
		},
	}
}

func KeysIterator[K comparable, V any](xs map[K]V) iterable.Iterator[K] {
	c := make(chan K)
	go func() {
		for k := range xs {
			c <- k
		}
		close(c)
	}()
	return &iterable.FunctionIterator[K]{
		F: func() *K {
			val, ok := <-c
			if ok {
				return &val
			}
			return nil
		},
	}
}
