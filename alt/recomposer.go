// Copyright (c) 2020, Peter Ohler, All rights reserved.

package alt

import (
	"encoding/json"
	"reflect"

	"github.com/ohler55/ojg"
)

// DefaultRecomposer provides a shared Recomposer. Note that this should not
// be shared across go routines unless all types that will be used are
// registered first. That can be done explicitly or with a warm up run.
var DefaultRecomposer = Recomposer{
	composers: map[string]*composer{},
}

// RecomposeFunc should build an object from data in a map returning the
// recomposed object or an error.
type RecomposeFunc func(map[string]any) (any, error)

// RecomposeAnyFunc should build an object from data in an any
// returning the recomposed object or an error.
type RecomposeAnyFunc func(any) (any, error)

// Recomposer is used to recompose simple data into structs.
type Recomposer struct {

	// CreateKey identifies the creation key in decomposed objects.
	CreateKey string

	composers map[string]*composer

	// NumConvMethod specifies the json.Number conversion method.
	NumConvMethod ojg.NumConvMethod
}

var jsonUnmarshalerType reflect.Type

func init() {
	jsonUnmarshalerType = reflect.TypeOf((*json.Unmarshaler)(nil)).Elem()

}

// RegisterComposer regsiters a composer function for a value type. A nil
// function will still register the default composer which uses reflection.
func (r *Recomposer) RegisterComposer(val any, fun RecomposeFunc) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterAnyComposer regsiters a composer function for a value type. A nil
// function will still register the default composer which uses reflection.
func (r *Recomposer) RegisterAnyComposer(val any, fun RecomposeAnyFunc) error {
	_ = "STUB: not implemented"
	return nil
}

// RegisterUnmarshalerComposer regsiters a composer function for a named
// value. This is only used to register cross package json.Unmarshaler
// composer which returns []byte.
func (r *Recomposer) RegisterUnmarshalerComposer(fun RecomposeAnyFunc) {
	_ = "STUB: not implemented"
	return
}

func (r *Recomposer) registerComposer(rt reflect.Type, fun RecomposeFunc, typeName string) (*composer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// TBD could loosen this up and allow any type as long as a function is provided. id is path through field names

// If already registered then there is no reason to walk the fields again.

// Private fields should be skipped.

func (r *Recomposer) registerAnyComposer(rt reflect.Type, fun RecomposeAnyFunc) (*composer, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Recompose simple data into more complex go types.
func (r *Recomposer) Recompose(v any, tv ...any) (out any, err error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// MustRecompose simple data into more complex go types.
func (r *Recomposer) MustRecompose(v any, tv ...any) (out any) {
	_ = "STUB: not implemented"
	return *new(any)
}

// Special case. Must return []byte.

func (r *Recomposer) recompAny(v any) any { _ = "STUB: not implemented"; return *new(any) }

// This small rounding makes the conversion from 32 bit to 64 bit
// display nicer.

func (r *Recomposer) recomp(v any, rv reflect.Value, typeName string) {
	_ = "STUB: not implemented"
	return
}

// Kind of awkward but the double reflect is needed to get the
// actual type of the element value if the slice input is []any.

func (r *Recomposer) setValue(v any, rv reflect.Value, sf *reflect.StructField, parent string) {
	_ = "STUB: not implemented"
	return
}

// Special case. Must return []byte.
