// Copyright (c) 2021, Peter Ohler, All rights reserved.

package alt

import (
	"reflect"
)

var int32ValFuncs = [8]valFunc{
	valInt32,
	valInt32AsString,
	valInt32NotEmpty,
	valInt32NotEmptyAsString,
	ivalInt32,
	ivalInt32AsString,
	ivalInt32NotEmpty,
	ivalInt32NotEmptyAsString,
}

func valInt32(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Value), false
}

func valInt32AsString(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Value), false
}

func valInt32NotEmpty(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Value), false
}

func valInt32NotEmptyAsString(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Value), false
}

func ivalInt32(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Value), false
}

func ivalInt32AsString(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Value), false
}

func ivalInt32NotEmpty(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Value), false
}

func ivalInt32NotEmptyAsString(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Value), false
}
