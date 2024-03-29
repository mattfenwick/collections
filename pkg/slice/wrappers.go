package slice

import (
	. "github.com/mattfenwick/collections/pkg/base"
)

func SliceEq[A Eq[A]](a []A) *EqBoxBy[[]A] {
	return BoxEqBy[[]A](a, EqualPairwiseEq[A]())
}

func SliceOrd[A Ord[A]](a []A) *OrdBoxBy[[]A] {
	return BoxOrdBy[[]A](a, EqualPairwiseEq[A](), ComparePairwiseBy(Compare[A]))
}

//type SliceEq[A Eq[A]] []A
//
//func (xs SliceEq[A]) Equal(ys SliceEq[A]) bool {
//	return EqualSlicePairwiseEq[A]()(xs, ys)
//}

//type SliceOrd[A Ord[A]] []A
//
//func (xs SliceOrd[A]) Compare(ys SliceOrd[A]) Ordering {
//	return ComparePairwiseBy(Compare[A])(xs, ys)
//}
//
//func (xs SliceOrd[A]) Equal(ys SliceOrd[A]) bool {
//	return EqualSlicePairwiseEq[A]()(xs, ys)
//}
