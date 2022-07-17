package main

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"reflect"
	"strconv"
)

type MySlice []int

func main() {
	a := &StuffA{Const: "mnop"}
	b := &StuffB{}
	c := &StuffC[int]{}

	fmt.Println(DoStuff(a, a))
	fmt.Println(DoStuff(b, b))
	//fmt.Println(DoStuff(a, b))
	//fmt.Println(DoStuff(b, a))

	fmt.Println(Example(a, a))
	fmt.Println(Example(c, c))

	g := &Generic[StringAlias, Getter[StringAlias], Setter[StringAlias]]{
		A: "abc",
	}
	g.B = &GetterImpl[StringAlias]{A: "qrs"} // g
	g.C = g
	fmt.Printf("%+v\n%s\n", g.Get(), g.GenericFunc())

	xs := MySlice{1, 2, 3, 4, 5, 6}
	doubled := Double(xs)
	doubledChanges := DoubleChangesType(xs)
	fmt.Printf("%+v  %s\n%+v  %s\n%s\n",
		doubled, reflect.TypeOf(doubled),
		doubledChanges, reflect.TypeOf(doubledChanges),
		reflect.TypeOf(xs))
}

type Stuff interface {
	DoStuff() string
}

type StuffA struct {
	Const string
}

func (s *StuffA) DoStuff() string {
	return s.Const
}

type StuffB struct{}

func (s *StuffB) DoStuff() string {
	return "abc"
}

type StuffC[X any] struct {
	FieldX X
}

func DoStuff[A Stuff](a A, b A) string {
	return fmt.Sprintf("%s: %s", a.DoStuff(), b.DoStuff())
}

func Example[A any](a A, b A) string {
	dumped, err := json.MarshalIndent([]interface{}{a, b}, "", "  ")
	DoOrDie(err)
	return string(dumped)
}

// StringAlias is a "type definition", NOT a type alias
//   a type definition can be used with an interface type constraint
//   i.e. "~string" , whereas without the "~",
//   this couldn't be used where a string is needed
type StringAlias string

type IntOrString interface {
	int | ~string
}

type Wrapper[T IntOrString] struct {
	T T
}

func (w *Wrapper[T]) GetInt() (int, error) {
	return 0, nil
	// looks like we can't do this
	//switch v := w.T.(type) {
	//case int:
	//	return v, nil
	//case string:
	//	return 0, errors.Errorf("can't get int -- have string")
	//default:
	//	panic("this is impossible")
	//}
}

func WrapperExample() {
	//w := &Wrapper[int]{T: 3}
	//w1 := &Wrapper[string]{T: "qrs"}
}

// Note: these functions cannot be written; compiler error:
//   interface contains type constraints
//func ReturnIntOrString() IntOrString {
//	return 3
//}
//func ConsumeIntOrString(iors IntOrString) {
//
//}
// this is illegal because of the T: "cannot embed a type parameter"
//type IllegalInterface[T any] interface {
//	int | uint | T
//}
// this is illegal because: "cannot use error in union (error contains methods)"
//type ErrorOrInt interface {
//	error | int
//}
// ErrorAndInt however, *is* legal:
type ErrorAndInt interface {
	int
	error
}

type Setter[A any] interface {
	Set(A)
}

type Getter[A any] interface {
	Get() A
}

type GetterImpl[A any] struct{ A A }

func (g *GetterImpl[A]) Get() A {
	return g.A
}

type Generic[A IntOrString, B Getter[A], C Setter[A]] struct {
	A A
	B B
	C C
}

func (g *Generic[A, B, C]) GenericFunc() string {
	one := g.B.Get()
	two := g.Get()

	g.C.Set(two)
	g.Set(one)

	return fmt.Sprintf("%+v\n%+v\n%+v\n%+v\n\n", one, two, g.B.Get(), g.Get())
}

func (g *Generic[A, B, C]) Get() A {
	return g.A
}

func (g *Generic[A, B, C]) Set(a A) {
	g.A = a
}

type Integer interface {
	uint | uint8 | uint16 | uint32 | uint64 | int | int8 | int16 | int32 | int64
}

// Double is kinda weird, this is how you get typedefs to get carried through to return values
//   see https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#element-constraint-example
//   compare to: Double[A Integer](xs []A) []A , where the return type is different
func Double[S ~[]A, A Integer](xs S) S {
	r := make([]A, len(xs))
	for i, v := range xs {
		r[i] = v + v
	}
	return r
}

func DoubleChangesType[A Integer](xs []A) []A {
	r := make([]A, len(xs))
	for i, v := range xs {
		r[i] = v + v
	}
	return r
}

// ExampleSetter is from: https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#pointer-method-example
type ExampleSetter interface {
	Set(string)
}

func FromStrings[T ExampleSetter](s []string) []T {
	result := make([]T, len(s))
	for i, v := range s {
		result[i].Set(v)
	}
	return result
}

type Settable int

func (p *Settable) Set(s string) {
	i, err := strconv.Atoi(s)
	DoOrDie(err)
	*p = Settable(i)
}

func SettableExample() {
	//nums := FromStrings[Settable]([]string{"1", "2"}) // doesn't compile
	nums := FromStrings[*Settable]([]string{"1", "2"})
	fmt.Printf("%s  %+v\n", reflect.TypeOf(nums), nums)
}

// By combining ~ constraints and method constraints, it may be possible to
//   to create a set of wrappers for built-in golang functionality that get
//   rid of the magic of builtins so they can be used in a principled, consistent way
//  see https://go.googlesource.com/proposal/+/refs/heads/master/design/43651-type-parameters.md#both-elements-and-methods-in-constraints

type byteseq interface {
	string | []byte
}

func Join[T byteseq](a T, sep T) {

}

func Eg() {
	a := "abc"
	b := []byte(a)
	Join(a, a)
	Join(b, b)
	//Join(a, b)
}

// TODO
//type Monad[M, A any] interface {
//	Join(m M[M[M, A], A])
//}
//
//func Bind[M Monad[M, A], A any](m M, f func(A) M) {}

func DoOrDie(err error) {
	if err != nil {
		logrus.Fatalf("%+v", err)
	}
}

//func Equalable[T comparable](a T, b T) bool {
//	return a == b
//}
//
//func GreaterThan[T comparable](a T, b T) bool {
//	return a > b
//}

//type A interface {
//	~int | ~uint
//}
//
//type B[Q A] interface {
//	Q
//}
//
//type Comparable[T comparable] struct {
//	T T
//}
//
//func (c *Comparable[T]) Equal(other *Comparable[T]) bool {
//	return c.T == other.T
//}
