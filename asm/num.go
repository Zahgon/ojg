// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "num?",
		Eval: num,
		Desc: `Returns true if the single required argumement is number
otherwise false is returned.`,
	})
}

func num(root map[string]any, at any, args ...any) any { _ = "STUB: not implemented"; return *new(any) }
