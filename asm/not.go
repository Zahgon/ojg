// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "not",
		Eval: not,
		Desc: `Returns the boolean NOT of the argument. Exactly one argument
is expected and it must be a boolean.`,
	})
}

func not(root map[string]any, at any, args ...any) any { _ = "STUB: not implemented"; return *new(any) }
