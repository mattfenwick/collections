package maybe

import . "github.com/mattfenwick/collections/pkg/base"

type Maybe[A any] struct {
	Value *A
}

func Just[A any](a A) *Maybe[A] {
	return &Maybe[A]{Value: &a}
}

func Nothing[A any]() *Maybe[A] {
	return &Maybe[A]{Value: nil}
}

func Map[A, B any](f F1[A, B], m *Maybe[A]) *Maybe[B] {
	if m.Value == nil {
		return Nothing[B]()
	}
	return Just(f(*m.Value))
}

func Reduce[A, B any](m *Maybe[A], b B, f F2[B, A, B]) B {
	if m.Value == nil {
		return b
	}
	return f(b, *m.Value)
}

func Filter[A any](m *Maybe[A], f F1[A, bool]) *Maybe[A] {
	if m.Value == nil {
		return m
	}
	if f(*m.Value) {
		return m
	}
	return Nothing[A]()
}

func Bind[A, B any](m *Maybe[A], f F1[A, *Maybe[B]]) *Maybe[B] {
	if m.Value == nil {
		return Nothing[B]()
	}
	return f(*m.Value)
}

func (m *Maybe[A]) Default(defaultValue A) A {
	if m.Value != nil {
		return *m.Value
	}
	return defaultValue
}
