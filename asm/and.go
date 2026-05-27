// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "and",
		Eval: and,
		Desc: `Returns true if all argument evaluate to true. Any arguments
that do not evaluate to a boolean or null (false) raise an error.`,
	})
}

func and(root map[string]any, at any, args ...any) any { _ = "STUB: not implemented"; return *new(any) }
