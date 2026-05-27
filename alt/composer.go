// Copyright (c) 2020, Peter Ohler, All rights reserved.

package alt

import (
	"reflect"
)

type composer struct {
	fun     RecomposeFunc
	any     RecomposeAnyFunc
	short   string
	full    string
	rtype   reflect.Type
	indexes map[string]reflect.StructField
}

func indexType(rt reflect.Type) (im map[string]reflect.StructField) {
	_ = "STUB: not implemented"
	return nil
}

// prepend index and add to im
