package function

func OnHelper[A, B, C any](combine func(B, B) C, project func(A) B, x A, y A) C {
	return combine(project(x), project(y))
}

func Cast[A any](val interface{}) A {
	return val.(A)
}

func First[A, B any](a A, b B) A {
	return a
}

func Second[A, B any](a A, b B) B {
	return b
}
