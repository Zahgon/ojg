// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "zone",
		Eval: zone,
		Desc: `Changes the timezone on a time to the location specified in the
second argument. Raises an error if the first argument does not
evaluate to a time or the location can not be determined.
Location can be either a string or the number of minutes offset
from UTC.`,
	})
}

func zone(root map[string]any, at any, args ...any) any {
	_ = "STUB: not implemented"
	return *new(any)
}
