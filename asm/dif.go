// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "dif",
		Eval: dif,
		Desc: `Returns the difference of all arguments. All arguments must be
numbers. If any of the arguments are not a number an error is
raised.`,
	})
	Define(&Fn{
		Name: "-",
		Eval: dif,
		Desc: `Returns the difference of all arguments. All arguments must be
numbers. If any of the arguments are not a number an error is
raised.`,
	})
}

func dif(root map[string]any, at any, args ...any) any { _ = "STUB: not implemented"; return *new(any) }
