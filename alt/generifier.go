// Copyright (c) 2020, Peter Ohler, All rights reserved.

package alt

import (
	"reflect"

	"github.com/ohler55/ojg/gen"
)

// Genericer is the interface for the Generic() function that converts types
// to generic types.
type Genericer interface {

	// Generic should return a Node that represents the object. Generally this
	// includes the use of a creation key consistent with call to the
	// reflection based Generic() function.
	Generic() gen.Node
}

// Generify converts a value into Node compliant data. A best effort is made
// to convert values that are not simple into generic Nodes.
func Generify(v any, options ...*Options) (n gen.Node) {
	_ = "STUB: not implemented"
	return *new(gen.Node)
}

// TBD OmitEmpty

// GenAlter converts a simple go data element into Node compliant data. A best
// effort is made to convert values that are not simple into generic Nodes. It
// modifies the values inplace if possible by altering the original.
func GenAlter(v any, options ...*Options) (n gen.Node) {
	_ = "STUB: not implemented"
	return *new(gen.Node)
}

// TBD OmitEmpty

// TBD delete in place

func reflectGenData(data any, opt *Options) gen.Node {
	_ = "STUB: not implemented"
	return *new(gen.Node)
}

func reflectGenValue(rv reflect.Value, opt *Options) (v gen.Node) {
	_ = "STUB: not implemented"
	return *new(gen.Node)
}

func reflectGenStruct(rv reflect.Value, opt *Options) gen.Node {
	_ = "STUB: not implemented"
	return *new(gen.Node)
}

// not a public field

// TBD OmitEmpty

func reflectGenComplex(rv reflect.Value, opt *Options) gen.Node {
	_ = "STUB: not implemented"
	return *new(gen.Node)
}

func reflectGenMap(rv reflect.Value, opt *Options) gen.Node {
	_ = "STUB: not implemented"
	return *new(gen.Node)
}

// TBD OmitEmpty

func reflectGenArray(rv reflect.Value, opt *Options) gen.Node {
	_ = "STUB: not implemented"
	return *new(gen.Node)
}
