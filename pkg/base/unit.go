package base

// Unit represents the Haskell value `()`
type Unit struct{}

var UnitC = &Unit{}

func (u *Unit) Equal(other *Unit) bool {
	return true
}

func (u *Unit) Compare(other *Unit) Ordering {
	return OrderingEqual
}
