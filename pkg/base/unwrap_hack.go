package base

type UW[A any, B any] interface {
	Unwrap() B
}

func (a Bool) Unwrap() bool {
	return bool(a)
}

func Yes[A any, B any](t UW[A, B]) B {
	return t.Unwrap()
}

func No(a Bool) bool {
	return Yes[Bool, bool](a)
}

//func Unwrap[A any, B any](f func(A) B, a A) B {
//	return f(a)
//}
//
//func UnwrapE(a Bool) bool {
//	return Unwrap[Bool, bool](func(b Bool) bool {
//		return bool(b)
//	}, a)
//}
