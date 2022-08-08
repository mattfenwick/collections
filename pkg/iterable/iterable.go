package iterable

import "context"

type Iterator[A any] interface {
	Next() *A // TODO should this be (A, bool) ?
}

type Iterable[A any] interface {
	Iterator() Iterator[A]
}

func ToSlice[A any](i Iterator[A]) []A {
	var out []A
	for {
		next := i.Next()
		if next == nil {
			break
		}
		out = append(out, *next)
	}
	return out
}

func ToChannel[A any](ctx context.Context, i Iterator[A]) <-chan A {
	channel := make(chan A)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
			}
			next := i.Next()
			if next == nil {
				break
			}
			channel <- *next
		}
	}()
	return channel
}

type FunctionIterator[A any] struct {
	F func() *A
}

func (f *FunctionIterator[A]) Next() *A {
	return f.F()
}

//func SetIterator
