// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "equal",
		Eval: equal,
		Desc: `Returns true if all the argument are equal. Aliases are eq, ==,
and equal.`,
	})
	Define(&Fn{
		Name: "eq",
		Eval: equal,
		Desc: `Returns true if all the argument are equal. Aliases are eq, ==,
and equal.`,
	})
	Define(&Fn{
		Name: "==",
		Eval: equal,
		Desc: `Returns true if all the argument are equal. Aliases are eq, ==,
and equal.`,
	})
}

func equal(root map[string]any, at any, args ...any) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func equalVals(v0, v1 any) (eq bool) { _ = "STUB: not implemented"; return false }

func asInt(v any) (i int64, ok bool) { _ = "STUB: not implemented"; return 0, false }

func asFloat(v any) (f float64, ok bool) { _ = "STUB: not implemented"; return 0, false }
