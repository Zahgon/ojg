// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "root",
		Eval: root,
		Desc: `Forms a path starting with @. The remaining string arguments are
joined with a '.' and parsed to form a jp.Expr.`,
	})
}

func root(root map[string]any, at any, args ...any) any {
	_ = "STUB: not implemented"
	return *new(any)
}
