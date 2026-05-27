// Copyright (c) 2021, Peter Ohler, All rights reserved.

package alt

import (
	"reflect"
)

var uintValFuncs = [8]valFunc{
	valUint,
	valUintAsString,
	valUintNotEmpty,
	valUintNotEmptyAsString,
	ivalUint,
	ivalUintAsString,
	ivalUintNotEmpty,
	ivalUintNotEmptyAsString,
}

func valUint(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Value), false
}

func valUintAsString(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Value), false
}

func valUintNotEmpty(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Value), false
}

func valUintNotEmptyAsString(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Value), false
}

func ivalUint(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Value), false
}

func ivalUintAsString(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Value), false
}

func ivalUintNotEmpty(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Value), false
}

func ivalUintNotEmptyAsString(fi *finfo, rv reflect.Value, addr uintptr) (any, reflect.Value, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(reflect.Value), false
}
