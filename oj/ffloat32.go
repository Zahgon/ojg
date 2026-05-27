// Copyright (c) 2021, Peter Ohler, All rights reserved.

package oj

import (
	"reflect"
)

var float32AppendFuncs = [8]appendFunc{
	appendFloat32,
	appendFloat32AsString,
	appendFloat32NotEmpty,
	appendFloat32NotEmptyAsString,
	iappendFloat32,
	iappendFloat32AsString,
	iappendFloat32NotEmpty,
	iappendFloat32NotEmptyAsString,
}

func appendFloat32(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, any, appendStatus) {
	_ = "STUB: not implemented"
	return nil, *new(any), *new(appendStatus)
}

func appendFloat32AsString(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, any, appendStatus) {
	_ = "STUB: not implemented"
	return nil, *new(any), *new(appendStatus)
}

func appendFloat32NotEmpty(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, any, appendStatus) {
	_ = "STUB: not implemented"
	return nil, *new(any), *new(appendStatus)
}

func appendFloat32NotEmptyAsString(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, any, appendStatus) {
	_ = "STUB: not implemented"
	return nil, *new(any), *new(appendStatus)
}

func iappendFloat32(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, any, appendStatus) {
	_ = "STUB: not implemented"
	return nil, *new(any), *new(appendStatus)
}

func iappendFloat32AsString(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, any, appendStatus) {
	_ = "STUB: not implemented"
	return nil, *new(any), *new(appendStatus)
}

func iappendFloat32NotEmpty(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, any, appendStatus) {
	_ = "STUB: not implemented"
	return nil, *new(any), *new(appendStatus)
}

func iappendFloat32NotEmptyAsString(fi *finfo, buf []byte, rv reflect.Value, addr uintptr, safe bool) ([]byte, any, appendStatus) {
	_ = "STUB: not implemented"
	return nil, *new(any), *new(appendStatus)
}
