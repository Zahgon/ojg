// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "substr",
		Eval: substr,
		Desc: `Returns a substring of the input string. The second argument
must be an integer that marks the start of the substring. The
third integer argument indicates the length of the substring
if provided. If the length argument is not provided the end of
the substring is the end of the input string.`,
	})
}

func substr(root map[string]any, at any, args ...any) any {
	_ = "STUB: not implemented"
	return *new(any)
}
