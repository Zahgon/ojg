// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "quotient",
		Eval: quotient,
		Desc: `Returns the quotient of all arguments. All arguments must be
numbers. If any of the arguments are not a number an error is
raised. If an attempt is made to divide by zero and error will
be raised.`,
	})
	Define(&Fn{
		Name: "/",
		Eval: quotient,
		Desc: `Returns the quotient of all arguments. All arguments must be
numbers. If any of the arguments are not a number an error is
raised. If an attempt is made to divide by zero and error will
be raised.`,
	})
}

func quotient(root map[string]any, at any, args ...any) any {
	_ = "STUB: not implemented"
	return *new(any)
}
