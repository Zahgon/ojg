// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "sum",
		Eval: sum,
		Desc: `Returns the sum of all arguments. All arguments must be numbers
or strings. If any argument is a string then the result will be
a string otherwise the result will be a number. If any of the
arguments are not a number or a string an error is raised.`,
	})
	Define(&Fn{
		Name: "+",
		Eval: sum,
		Desc: `Returns the sum of all arguments. All arguments must be numbers
or strings. If any argument is a string then the result will be
a string otherwise the result will be a number. If any of the
arguments are not a number or a string an error is raised.`,
	})
}

const (
	intSum = iota
	floatSum
	strSum
)

func sum(root map[string]any, at any, args ...any) any { _ = "STUB: not implemented"; return *new(any) }
