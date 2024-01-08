package slice

import (
	. "github.com/mattfenwick/collections/pkg/base"
)

type SliceEq[A Eq[A]] []A

func (xs SliceEq[A]) Equal(ys SliceEq[A]) bool {
	return EqualSlicePairwiseEq[A]()(xs, ys)
}

type SliceOrd[A Ord[A]] []A

func (xs SliceOrd[A]) Compare(ys SliceOrd[A]) Ordering {
	return CompareSlicePairwiseBy(Compare[A])(xs, ys)
}

func (xs SliceOrd[A]) Equal(ys SliceOrd[A]) bool {
	return EqualSlicePairwiseEq[A]()(xs, ys)
}
