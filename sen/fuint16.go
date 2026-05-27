// Copyright (c) 2021, Peter Ohler, All rights reserved.

package sen

import (
	"reflect"
)

var uint16AppendFuncs = [8]appendFunc{
	appendUint16,
	appendUint16AsString,
	appendUint16NotEmpty,
	appendUint16NotEmptyAsString,
	iappendUint16,
	iappendUint16AsString,
	iappendUint16NotEmpty,
	iappendUint16NotEmptyAsString,
}

func appendUint16(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, any, appendStatus) {
	_ = "STUB: not implemented"
	return nil, *new(any), *new(appendStatus)
}

func appendUint16AsString(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, any, appendStatus) {
	_ = "STUB: not implemented"
	return nil, *new(any), *new(appendStatus)
}

func appendUint16NotEmpty(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, any, appendStatus) {
	_ = "STUB: not implemented"
	return nil, *new(any), *new(appendStatus)
}

func appendUint16NotEmptyAsString(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, any, appendStatus) {
	_ = "STUB: not implemented"
	return nil, *new(any), *new(appendStatus)
}

func iappendUint16(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, any, appendStatus) {
	_ = "STUB: not implemented"
	return nil, *new(any), *new(appendStatus)
}

func iappendUint16AsString(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, any, appendStatus) {
	_ = "STUB: not implemented"
	return nil, *new(any), *new(appendStatus)
}

func iappendUint16NotEmpty(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, any, appendStatus) {
	_ = "STUB: not implemented"
	return nil, *new(any), *new(appendStatus)
}

func iappendUint16NotEmptyAsString(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, any, appendStatus) {
	_ = "STUB: not implemented"
	return nil, *new(any), *new(appendStatus)
}
