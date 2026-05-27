// Copyright (c) 2021, Peter Ohler, All rights reserved.

package alt

import (
	"reflect"
)

var uint8ValFuncs = [8]valFunc{
	valUint8,
	valUint8AsString,
	valUint8NotEmpty,
	valUint8NotEmptyAsString,
	ivalUint8,
	ivalUint8AsString,
	ivalUint8NotEmpty,
	ivalUint8NotEmptyAsString,
}

func valUint8(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Value), false
}

func valUint8AsString(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Value), false
}

func valUint8NotEmpty(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Value), false
}

func valUint8NotEmptyAsString(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Value), false
}

func ivalUint8(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Value), false
}

func ivalUint8AsString(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Value), false
}

func ivalUint8NotEmpty(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Value), false
}

func ivalUint8NotEmptyAsString(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Value), false
}
