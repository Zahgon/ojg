// Copyright (c) 2021, Peter Ohler, All rights reserved.

package sen

import (
	"reflect"
)

var int32AppendFuncs = [8]appendFunc{
	appendInt32,
	appendInt32AsString,
	appendInt32NotEmpty,
	appendInt32NotEmptyAsString,
	iappendInt32,
	iappendInt32AsString,
	iappendInt32NotEmpty,
	iappendInt32NotEmptyAsString,
}

func appendInt32(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, any, appendStatus) {
	_ = "STUB: not implemented"
	return nil, *new(any), *new(appendStatus)
}

func appendInt32AsString(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, any, appendStatus) {
	_ = "STUB: not implemented"
	return nil, *new(any), *new(appendStatus)
}

func appendInt32NotEmpty(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, any, appendStatus) {
	_ = "STUB: not implemented"
	return nil, *new(any), *new(appendStatus)
}

func appendInt32NotEmptyAsString(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, any, appendStatus) {
	_ = "STUB: not implemented"
	return nil, *new(any), *new(appendStatus)
}

func iappendInt32(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, any, appendStatus) {
	_ = "STUB: not implemented"
	return nil, *new(any), *new(appendStatus)
}

func iappendInt32AsString(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, any, appendStatus) {
	_ = "STUB: not implemented"
	return nil, *new(any), *new(appendStatus)
}

func iappendInt32NotEmpty(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, any, appendStatus) {
	_ = "STUB: not implemented"
	return nil, *new(any), *new(appendStatus)
}

func iappendInt32NotEmptyAsString(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, any, appendStatus) {
	_ = "STUB: not implemented"
	return nil, *new(any), *new(appendStatus)
}
