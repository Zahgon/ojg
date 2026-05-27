// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "null?",
		Eval: null,
		Desc: `Returns true if the single required argumement is null (JSON)
or nil (golang) otherwise false is returned.`,
	})
	Define(&Fn{
		Name: "nil?",
		Eval: null,
		Desc: `Returns true if the single required argumement is null (JSON)
or nil (golang) otherwise false is returned.`,
	})
}

func null(root map[string]any, at any, args ...any) any {
	_ = "STUB: not implemented"
	return *new(any)
}
