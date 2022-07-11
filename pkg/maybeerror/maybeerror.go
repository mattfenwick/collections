package maybeerror

type Failure struct{}

var FailureC = &Failure{}

type Success[A any] struct {
	Value A
}

type Error[E any] struct {
	Value E
}

type MaybeError[E any, A any] struct {
	Success *Success[A]
	Failure *Failure
	Error   *Error[E]
}

func NewSuccess[E any, A any](value A) *MaybeError[E, A] {
	return &MaybeError[E, A]{Success: &Success[A]{Value: value}}
}

func NewError[E any, A any](error E) *MaybeError[E, A] {
	return &MaybeError[E, A]{Error: &Error[E]{Value: error}}
}

func NewFailure[E any, A any]() *MaybeError[E, A] {
	return &MaybeError[E, A]{Failure: FailureC}
}

// IsValid is a debugging check
func (m *MaybeError[E, A]) IsValid() bool {
	present := 0
	for _, v := range []interface{}{m.Success, m.Failure, m.Error} {
		if v != nil {
			present++
		}
	}
	return present == 1
}

func Fmap[E, A, B any](m *MaybeError[E, A], f func(a A) B) *MaybeError[E, B] {
	if m.Success != nil {
		return NewSuccess[E, B](f(m.Success.Value))
	} else if m.Failure != nil {
		return NewFailure[E, B]()
	}
	return NewError[E, B](m.Error.Value)
}

func App2[E, A, B, C any](f func(A, B) C, m1 *MaybeError[E, A], m2 *MaybeError[E, B]) *MaybeError[E, C] {
	return Bind[E, A, C](m1, func(a A) *MaybeError[E, C] {
		return Bind[E, B, C](m2, func(b B) *MaybeError[E, C] {
			return NewSuccess[E, C](f(a, b))
		})
	})
}

func Bind[E, A, B any](m *MaybeError[E, A], f func(A) *MaybeError[E, B]) *MaybeError[E, B] {
	if m.Success != nil {
		return f(m.Success.Value)
	} else if m.Failure != nil {
		return NewFailure[E, B]()
	}
	return NewError[E, B](m.Error.Value)
}

func MapError[E1, E2, A any](m *MaybeError[E1, A], f func(E1) E2) *MaybeError[E2, A] {
	if m.Success != nil {
		return NewSuccess[E2, A](m.Success.Value)
	} else if m.Failure != nil {
		return NewFailure[E2, A]()
	}
	return NewError[E2, A](f(m.Error.Value))
}

func (m *MaybeError[E, A]) Plus(other *MaybeError[E, A]) *MaybeError[E, A] {
	if m.Success != nil {
		return m
	} else if m.Failure != nil {
		return other
	}
	return m
}
