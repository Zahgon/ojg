// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "product",
		Eval: product,
		Desc: `Returns the product of all arguments. All arguments must be
numbers. If any of the arguments are not a number an error is
raised.`,
	})
	Define(&Fn{
		Name: "*",
		Eval: product,
		Desc: `Returns the product of all arguments. All arguments must be
numbers. If any of the arguments are not a number an error is
raised.`,
	})
}

func product(root map[string]any, at any, args ...any) any {
	_ = "STUB: not implemented"
	return *new(any)
}
