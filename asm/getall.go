// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "getall",
		Eval: getall,
		Desc: `Gets all matching values in either the root ($), or local (@),
or if present, the second argument. The required first argument
must be a path and the option second argument is the
data to apply the path to. The jp.Get() function is used to get
the results`,
	})
}

func getall(root map[string]any, at any, args ...any) (list any) {
	_ = "STUB: not implemented"
	return *new(any)
}
