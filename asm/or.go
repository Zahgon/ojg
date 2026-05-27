// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "or",
		Eval: or,
		Desc: `Returns true if any of the argument evaluate to true. Any
arguments that do not evaluate to a boolean or null (false)
raise an error.`,
	})
}

func or(root map[string]any, at any, args ...any) any { _ = "STUB: not implemented"; return *new(any) }
