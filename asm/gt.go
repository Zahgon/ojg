// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "gt",
		Eval: gt,
		Desc: `Returns true if each argument is greater than any subsequent
argument. An alias is >.`,
	})
	Define(&Fn{
		Name: ">",
		Eval: gt,
		Desc: `Returns true if each argument is greater than any subsequent
argument. An alias is gt.`,
	})
}

func gt(root map[string]any, at any, args ...any) any { _ = "STUB: not implemented"; return *new(any) }
