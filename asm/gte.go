// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "gte",
		Eval: gte,
		Desc: `Returns true if each argument is greater than or equal to any
subsequent argument. An alias is >=.`,
	})
	Define(&Fn{
		Name: ">=",
		Eval: gte,
		Desc: `Returns true if each argument is greater than or equal to any
subsequent argument. An alias is gte.`,
	})
}

func gte(root map[string]any, at any, args ...any) any { _ = "STUB: not implemented"; return *new(any) }
