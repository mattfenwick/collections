package builtin

import "golang.org/x/exp/constraints"

// Number is built out of:
//
//	https://pkg.go.dev/golang.org/x/exp@v0.0.0-20220706164943-b4a6d9510983/constraints
type Number interface {
	constraints.Integer | constraints.Float
}
