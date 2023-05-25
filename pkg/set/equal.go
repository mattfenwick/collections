package set

//import (
//	"github.com/mattfenwick/collections/pkg/base"
//	"github.com/mattfenwick/collections/pkg/builtin"
//	"github.com/mattfenwick/collections/pkg/slice"
//	"golang.org/x/exp/constraints"
//)
//
//func Equal[A constraints.Ordered]() base.Equaler[*Set[A]] {
//	return EqualBy(builtin.Equal[A])
//}
//
//func EqualBy[A any](equal base.Equaler[A]) base.Equaler[*Set[A]] {
//	return func(a *Set[A], b *Set[A]) bool {
//		return slice.EqualSlicePairwiseBy(equal)(slice.Sort(a.ToSlice()), slice.Sort(b.ToSlice()))
//	}
//}
