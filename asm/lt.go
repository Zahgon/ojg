// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "lt",
		Eval: lt,
		Desc: `Returns true if each argument is less than any subsequent
argument. An alias is <.`,
	})
	Define(&Fn{
		Name: "<",
		Eval: lt,
		Desc: `Returns true if each argument is less than any subsequent
argument. An alias is lt.`,
	})
}

func lt(root map[string]any, at any, args ...any) any { _ = "STUB: not implemented"; return *new(any) }
