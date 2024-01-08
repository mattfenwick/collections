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


### Method can't have type parameters

Example: you have a type `type Table[A any, V any] struct`

You'd like to add a method to map values, such as:

```golang
func (t *Table[A, V])MapValues[A any, V any, W any](f func(V) W) *Table[A, W] {
```

However, this won't compile -- `method can't have type parameters`

Practical implications: some functionality has to be implemented as free functions instead of
methods, which is a bit annoying in terms of autocomplete and findability of functionality.
