// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "mod",
		Eval: mod,
		Desc: `Returns the remainer of a modulo operation on the first two
argument. Both arguments must be integers and are both required.
An error is raised if the wrong argument types are given.`,
	})
}

func mod(root map[string]any, at any, args ...any) any { _ = "STUB: not implemented"; return *new(any) }
