// Copyright (c) 2020, Peter Ohler, All rights reserved.

package alt

import (
	"reflect"
)

// 23 for fraction in IEEE 754 which amounts to 7 significant digits. Use base
// 10 so that numbers look correct when displayed in base 10.
const fracMax = 10000000.0

func decompose(v any, opt *Options) any { _ = "STUB: not implemented"; return *new(any) }

// This small rounding makes the conversion from 32 bit to 64 bit
// display nicer.

func alter(v any, opt *Options) any { _ = "STUB: not implemented"; return *new(any) }

// This small rounding makes the conversion from 32 bit to 64 bit
// display nicer.

func reflectValue(rv reflect.Value, val any, opt *Options) (v any) {
	_ = "STUB: not implemented"
	return *new(any)
}

func reflectStruct(rv reflect.Value, val any, opt *Options) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func reflectEmbed(rv reflect.Value, val any, opt *Options) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func reflectComplex(rv reflect.Value, opt *Options) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func reflectMap(rv reflect.Value, opt *Options) any { _ = "STUB: not implemented"; return *new(any) }

func reflectArray(rv reflect.Value, opt *Options) any { _ = "STUB: not implemented"; return *new(any) }

func isNil(rv reflect.Value) bool { _ = "STUB: not implemented"; return false }

func condMapSet(m map[string]any, key string, value any, opt *Options) {
	_ = "STUB: not implemented"
	return
}
