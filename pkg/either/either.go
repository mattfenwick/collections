package either

type Either[E any, A any] struct {
	Success *A
	Error   *E
}

func Success[E any, A any](value A) *Either[E, A] {
	return &Either[E, A]{Success: &value}
}

func Error[E any, A any](error E) *Either[E, A] {
	return &Either[E, A]{Error: &error}
}

// IsValid is a debugging check
func (e *Either[E, A]) IsValid() bool {
	present := 0
	for _, v := range []interface{}{e.Success, e.Error} {
		if v != nil {
			present++
		}
	}
	return present == 1
}

func Map[E, A, B any](m *Either[E, A], f func(a A) B) *Either[E, B] {
	if m.Success != nil {
		return Success[E, B](f(*m.Success))
	}
	return Error[E, B](*m.Error)
}

func App2[E, A, B, C any](f func(A, B) C, m1 *Either[E, A], m2 *Either[E, B]) *Either[E, C] {
	return Bind[E, A, C](m1, func(a A) *Either[E, C] {
		return Bind[E, B, C](m2, func(b B) *Either[E, C] {
			return Success[E, C](f(a, b))
		})
	})
}

func Bind[E, A, B any](m *Either[E, A], f func(A) *Either[E, B]) *Either[E, B] {
	if m.Success != nil {
		return f(*m.Success)
	}
	return Error[E, B](*m.Error)
}

func MapError[E1, E2, A any](m *Either[E1, A], f func(E1) E2) *Either[E2, A] {
	if m.Success != nil {
		return Success[E2, A](*m.Success)
	}
	return Error[E2, A](f(*m.Error))
}

func (e *Either[E, A]) Plus(other *Either[E, A]) *Either[E, A] {
	if e.Success != nil {
		return e
	}
	return other
}
