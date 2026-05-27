// Copyright (c) 2021, Peter Ohler, All rights reserved.

package alt

import (
	"reflect"
)

const (
	strMask   = byte(0x01)
	omitMask  = byte(0x02)
	embedMask = byte(0x04)
)

var nilValue reflect.Value

type valFunc func(fi *finfo, rv reflect.Value, addr uintptr) (v any, fv reflect.Value, omit bool)

type finfo struct {
	rt     reflect.Type
	key    string
	value  valFunc
	ivalue valFunc
	index  []int
	offset uintptr
}

func valString(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Value), false
}

func valStringNotEmpty(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Value), false
}

func valJustVal(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Value), false
}

func valPtrNotEmpty(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Value), false
}

func valSliceNotEmpty(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Value), false
}

func valSimplifier(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Value), false
}

func valSimplifierAddr(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Value), false
}

func valGenericer(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Value), false
}

func valGenericerAddr(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Value), false
}

func newFinfo(f *reflect.StructField, key string, fx byte) *finfo {
	_ = "STUB: not implemented"
	return nil
}

// replace as necessary later
// replace as necessary later

// Check for interfaces first since almost any type can implement one of
// the supported interfaces.
