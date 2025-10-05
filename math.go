// Package inside test module while learning go
package gomodulespackagesimports

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Adds two numbers, and returns the result
//
// [In-depth explanation of this novel functionality]
//
// [In-depth explanation of this novel functionality]: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}
