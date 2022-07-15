package base

type SliceEq[A Eq[A]] []A
type SliceOrd[A Ord[A]] []A

// EqOrComparable allows us to avoid getting "invalid map key type A (missing comparable constraint)"
//   errors, if we just used Eq[A] without this additional interface
type EqOrComparable[A Eq[A]] interface {
	Eq[A]
	comparable
}

type MapEq[A EqOrComparable[A], B Eq[B]] map[A]B
