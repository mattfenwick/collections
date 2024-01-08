# collections

Golang collections, generics, and utilities

# Create a new version

```bash
# figure out what version to use
git tag --list

# make a tag
git tag -a "$NEW_VERSION"

git push --tags
```

## Examples

See: [./cmd](./cmd)

## Notes

### Pointer types are `comparable` -- in a surprising way

This refers to the golang-builtin interface type parameter constraint.

The `comparable` definition says:

```golang
// comparable is an interface that is implemented by all comparable types
// (booleans, numbers, strings, pointers, channels, arrays of comparable types,
// structs whose fields are all comparable types).
// The comparable interface may only be used as a type parameter constraint,
// not as the type of a variable.
type comparable interface{ comparable }
```

Be careful if you: have two different objects with equal values.  These aren't equal.
For example:
```golang
    p1 := NewPair(1, 4)
    p2 := NewPair(1, 4)

    gomega.Expect([]bool{
        p1 == p2,
        p2 == p1,
        p1 == p1,
        p2 == p2,
    }).To(gomega.Equal([]bool{
        false, // surprise!
        false, // surprise!
        true,
        true,
    }))
```

Practical implications -- be careful when:

 - using `comparable` types as map keys
 - checking `comparable` types for equality
 - using `comparable` types in sets

### Cases where you'd like to use methods, but instead have to use free functions

Practical implications: some functionality has to be implemented as free functions instead of
methods, which is a bit annoying in terms of autocomplete and findability of functionality.


#### Method can't have type parameters

Example: you have a type `type Table[A any, V any] struct`

You'd like to add a method to map values, such as:

```golang
func (t *Table[A, V])MapValues[A any, V any, W any](f func(V) W) *Table[A, W] {
```

However, this won't compile -- `method can't have type parameters`


#### Method can't have additional type constraints

It's not possible to equip Pair with an `Eq` instance:

```golang
//func (p *Pair[A, B]) Equal[A Eq[A], B Eq[B]](p2 *Pair[A, B]) bool {
//	return p.Fst.Equal(p2.Fst)
//}
```

Note that we only want to have this instance available when *both* of the
type parameters are also `Eq` -- otherwise, this instance shouldn't be available.

Instead, we have to use a free function to bring in the type constraints:

```golang
func EqualPairEq[A Eq[A], B Eq[B]]() Equaler[*Pair[A, B]] {
	return EqualPairBy(Equal[A], Equal[B])
}}
```


## Open questions

### Interface embedding vs curiously-recurring type constraint

Related: [some background on CRTP](https://en.wikipedia.org/wiki/Curiously_recurring_template_pattern).

How, precisely, are these two examples different from each other and when should one or the
other be used?

```golang
type Ord[T any] interface {
  Eq[T]
  Compare(T) Ordering
}
```

```golang
type Ord2[T Eq[T]] interface {
    Compare(T) Ordering
}
```