// Copyright (c) 2021, Peter Ohler, All rights reserved.

package sen

import (
	"reflect"
)

const (
	strMask   = byte(0x01)
	omitMask  = byte(0x02)
	embedMask = byte(0x04)

	aJustKey appendStatus = iota
	aWrote
	aSkip
	aChanged
)

type appendStatus byte

type appendFunc func(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, any, appendStatus)

// Field hold information about a struct field.
type finfo struct {
	rt      reflect.Type
	key     string
	kind    reflect.Kind
	elem    *sinfo
	Append  appendFunc
	iAppend appendFunc
	jkey    []byte
	index   []int
	offset  uintptr
}

func (f *finfo) keyLen() int { _ = "STUB: not implemented"; return 0 }

func appendJustKey(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, any, appendStatus) {
	_ = "STUB: not implemented"
	return nil, *new(any), *new(appendStatus)
}

func appendPtrNotEmpty(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, any, appendStatus) {
	_ = "STUB: not implemented"
	return nil, *new(any), *new(appendStatus)
}

// real nil check

func appendSliceNotEmpty(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, any, appendStatus) {
	_ = "STUB: not implemented"
	return nil, *new(any), *new(appendStatus)
}

func appendSENString(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, any, appendStatus) {
	_ = "STUB: not implemented"
	return nil, *new(any), *new(appendStatus)
}

func appendSENStringNotEmpty(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, any, appendStatus) {
	_ = "STUB: not implemented"
	return nil, *new(any), *new(appendStatus)
}

func whichAppend(rt reflect.Type, omitEmpty bool) (f appendFunc) {
	_ = "STUB: not implemented"
	return *new(appendFunc)
}

func newFinfo(f *reflect.StructField, key string, omitEmpty, asString, pretty, embedded bool) *finfo {
	_ = "STUB: not implemented"
	return nil
}

// Check for interfaces first since almost any type can implement one of
// the supported interfaces.
