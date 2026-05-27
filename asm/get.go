// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "get",
		Eval: get,
		Desc: `Gets the first matching value in either the root ($), local (@),
or if present, the second argument. The required first argument
must be a path and the option second argument is the
data to apply the path to. The jp.First() function is used to
get the results`,
	})
}

func get(root map[string]any, at any, args ...any) (val any) {
	_ = "STUB: not implemented"
	return *new(any)
}
