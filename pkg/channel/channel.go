package channel

import "github.com/mattfenwick/collections/pkg/iterable"

type ReadOnlyChannel[A any] <-chan A

func (r ReadOnlyChannel[A]) Iterator() iterable.Iterator[A] {
	return Iterator(r)
}

func Iterator[A any](c <-chan A) iterable.Iterator[A] {
	return &iterable.FunctionIterator[A]{
		F: func() *A {
			value, ok := <-c
			if ok {
				return &value
			}
			return nil
		},
	}
}

type Channel[A any] chan A

type WriteOnlyChannel[A any] chan<- A
