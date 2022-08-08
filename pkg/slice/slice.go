package slice

import "github.com/mattfenwick/collections/pkg/iterable"

type Slice[A any] []A

func (s Slice[A]) Iterator() iterable.Iterator[A] {
	return Iterator(s)
}

func Iterator[A any](xs []A) iterable.Iterator[A] {
	i := 0
	return &iterable.FunctionIterator[A]{
		F: func() *A {
			if i < len(xs) {
				value := xs[i]
				i++
				return &value
			}
			return nil
		},
	}
}
