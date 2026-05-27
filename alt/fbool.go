// Copyright (c) 2021, Peter Ohler, All rights reserved.

package alt

import (
	"reflect"
)

var boolValFuncs = [8]valFunc{
	valBool,
	valBoolAsString,
	valBoolNotEmpty,
	valBoolNotEmptyAsString,
	ivalBool,
	ivalBoolAsString,
	ivalBoolNotEmpty,
	ivalBoolNotEmptyAsString,
}

func valBool(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Value), false
}

func valBoolAsString(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Value), false
}

func valBoolNotEmpty(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Value), false
}

func valBoolNotEmptyAsString(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Value), false
}

func ivalBool(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Value), false
}

func ivalBoolAsString(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Value), false
}

func ivalBoolNotEmpty(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Value), false
}

func ivalBoolNotEmptyAsString(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Value), false
}
