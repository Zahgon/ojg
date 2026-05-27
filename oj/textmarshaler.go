// Copyright (c) 2021, Peter Ohler, All rights reserved.

package oj

import (
	"reflect"
)

func appendTextMarshaler(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, any, appendStatus) {
	_ = "STUB: not implemented"
	return nil, *new(any), *new(appendStatus)
}

// real nil check

func appendTextMarshalerAddr(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, any, appendStatus) {
	_ = "STUB: not implemented"
	return nil, *new(any), *new(appendStatus)
}

func appendTextMarshalerNotEmpty(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, any, appendStatus) {
	_ = "STUB: not implemented"
	return nil, *new(any), *new(appendStatus)
}

// real nil check

func appendTextMarshalerVal(buf []byte, v any, safe bool) ([]byte, any, appendStatus) {
	_ = "STUB: not implemented"
	return nil, *new(any), *new(appendStatus)
}
