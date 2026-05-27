// Copyright (c) 2021, Peter Ohler, All rights reserved.

package alt

import (
	"github.com/ohler55/ojg"
)

// Options is an alias for ojg.Options
type Options = ojg.Options

// Converter is an alias for ojg.Converter
type Converter = ojg.Converter

var (
	// DefaultOptions are the default options for the this package.
	DefaultOptions = ojg.DefaultOptions
	// BrightOptions are the bright color options.
	BrightOptions = ojg.BrightOptions
	// GoOptions are the options that match the go json.Marshal behavior.
	GoOptions = ojg.GoOptions
	// HTMLOptions are the options that can be used to encode as HTML JSON.
	HTMLOptions = ojg.HTMLOptions

	// TimeRFC3339Converter converts RFC3339 string into time.Time when
	// parsing.
	TimeRFC3339Converter = ojg.TimeRFC3339Converter
	// TimeNanoConverter converts integer values to time.Time assuming the
	// integer are nonoseconds,
	TimeNanoConverter = ojg.TimeNanoConverter
	// MongoConverter converts mongodb decorations into the correct times.
	MongoConverter = ojg.MongoConverter
)

func init() {
	// Use different defaults for decompose except the Go defaults. Set
	// OmitNil and provide a CreateKey for all.
	DefaultOptions.OmitNil = true
	DefaultOptions.CreateKey = "type"
	BrightOptions.OmitNil = true
	BrightOptions.CreateKey = "type"
	HTMLOptions.OmitNil = true
	HTMLOptions.CreateKey = "type"
}

// Dup is an alias for Decompose.
func Dup(v any, options ...*ojg.Options) any { _ = "STUB: not implemented"; return *new(any) }

// Decompose creates a simple type converting non simple to simple types using
// either the Simplify() interface or reflection. Unlike Alter() a deep copy
// is returned leaving the original data unchanged.
func Decompose(v any, options ...*ojg.Options) any { _ = "STUB: not implemented"; return *new(any) }

// Alter the data into all simple types converting non simple to simple types
// using either the Simplify() interface or reflection. Unlike Decompose() map
// and slice members are modified if necessary to assure all elements are
// simple types.
func Alter(v any, options ...*ojg.Options) any { _ = "STUB: not implemented"; return *new(any) }

// Recompose simple data into more complex go types.
func Recompose(v any, tv ...any) (out any, err error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// MustRecompose simple data into more complex go types and panics on error.
func MustRecompose(v any, tv ...any) (out any) { _ = "STUB: not implemented"; return *new(any) }

// NewRecomposer creates a new instance. The composers are a map of objects
// expected and functions to recompose them. If no function is provided then
// reflection is used instead.
func NewRecomposer(
	createKey string,
	composers map[any]RecomposeFunc,
	anyComposers ...map[any]RecomposeAnyFunc) (rec *Recomposer, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MustNewRecomposer creates a new instance. The composers are a map of objects
// expected and functions to recompose them. If no function is provided then
// reflection is used instead. Panics on error.
func MustNewRecomposer(
	createKey string,
	composers map[any]RecomposeFunc,
	anyComposers ...map[any]RecomposeAnyFunc) *Recomposer {
	_ = "STUB: not implemented"
	return nil
}
