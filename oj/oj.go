// Copyright (c) 2020, Peter Ohler, All rights reserved.

package oj

import (
	"io"
	"sync"

	"github.com/ohler55/ojg"
	"github.com/ohler55/ojg/alt"
	"github.com/ohler55/ojg/jp"
)

// Options is an alias for ojg.Options
type Options = ojg.Options

// Builder is an aliase for alt.Builder.
type Builder = alt.Builder

var (
	// DefaultOptions are the default options for the this package.
	DefaultOptions = ojg.DefaultOptions
	// BrightOptions are the bright color options.
	BrightOptions = ojg.BrightOptions

	// HTMLOptions are the options that can be used to encode as HTML JSON.
	HTMLOptions = ojg.HTMLOptions

	goOptions  = ojg.GoOptions
	writerPool = sync.Pool{
		New: func() any {
			return &Writer{Options: DefaultOptions, buf: make([]byte, 0, 1024)}
		},
	}
	marshalPool = sync.Pool{
		New: func() any {
			return &Writer{Options: goOptions, buf: make([]byte, 0, 1024), strict: true}
		},
	}
	parserPool = sync.Pool{
		New: func() any {
			return &Parser{}
		},
	}
)

// Parse JSON into a simple type. Arguments are optional and can be a bool,
// func(any) bool for callbacks, or a chan any for chan based
// result delivery.
//
// A bool indicates the NoComment parser attribute should be set to the bool
// value.
//
// A func argument is the callback for the parser if processing multiple
// JSONs. If no callback function is provided the processing is limited to
// only one JSON.
//
// A chan argument will be used to deliver parse results.
func Parse(b []byte, args ...any) (n any, err error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// MustParse JSON into a simple type. Arguments are optional and can be a bool,
// func(any) bool for callbacks, or a chan any for chan based
// result delivery. Panics on error
//
// A bool indicates the NoComment parser attribute should be set to the bool
// value.
//
// A func argument is the callback for the parser if processing multiple
// JSONs. If no callback function is provided the processing is limited to
// only one JSON.
//
// A chan argument will be used to deliver parse results.
func MustParse(b []byte, args ...any) (n any) { _ = "STUB: not implemented"; return *new(any) }

// ParseString is similar to Parse except it takes a string
// argument to be parsed instead of a []byte.
func ParseString(s string, args ...any) (n any, err error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// MustParseString is similar to MustParse except it takes a string
// argument to be parsed instead of a []byte.
func MustParseString(s string, args ...any) (n any) { _ = "STUB: not implemented"; return *new(any) }

// Load a JSON from a io.Reader into a simple type. An error is returned
// if not valid JSON.
func Load(r io.Reader, args ...any) (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

// MustLoad a JSON from a io.Reader into a simple type. Panics on error.
func MustLoad(r io.Reader, args ...any) (n any) { _ = "STUB: not implemented"; return *new(any) }

// Validate a JSON string. An error is returned if not valid JSON.
func Validate(b []byte) error { _ = "STUB: not implemented"; return nil }

// ValidateString a JSON string. An error is returned if not valid JSON.
func ValidateString(s string) error { _ = "STUB: not implemented"; return nil }

// ValidateReader a JSON stream. An error is returned if not valid JSON.
func ValidateReader(r io.Reader) error { _ = "STUB: not implemented"; return nil }

// Unmarshal parses the provided JSON and stores the result in the value
// pointed to by vp.
func Unmarshal(data []byte, vp any, recomposer ...*alt.Recomposer) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// JSON returns a JSON string for the data provided. The data can be a
// simple type of nil, bool, int, floats, time.Time, []any, or
// map[string]any or a Node type, The args, if supplied can be an
// int as an indent or a *Options.
func JSON(data any, args ...any) string { _ = "STUB: not implemented"; return "" }

// Marshal returns a JSON string for the data provided. The data can be a
// simple type of nil, bool, int, floats, time.Time, []any, or
// map[string]any or a gen.Node type, The args, if supplied can be an
// int as an indent, *ojg.Options, or a *Writer. An error will be returned if
// the Option.Strict flag is true and a value is encountered that can not be
// encoded other than by using the %v format of the fmt package.
func Marshal(data any, args ...any) (out []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Write a JSON string for the data provided. The data can be a simple type of
// nil, bool, int, floats, time.Time, []any, or map[string]any
// or a Node type, The args, if supplied can be an int as an indent or a
// *Options.
func Write(w io.Writer, data any, args ...any) (err error) { _ = "STUB: not implemented"; return nil }

func pickWriter(arg any, strict bool) (wr *Writer) { _ = "STUB: not implemented"; return nil }

// Match parses a JSON document and calls onData when a data element that
// matches the target path is encountered.
func Match(data []byte, onData func(path jp.Expr, data any), targets ...jp.Expr) error {
	_ = "STUB: not implemented"
	return nil
}

// MatchString parses a JSON document and calls onData when a data element that
// matches the target path is encountered.
func MatchString(data string, onData func(path jp.Expr, data any), targets ...jp.Expr) error {
	_ = "STUB: not implemented"
	return nil
}

// MatchLoad parses a JSON document from an io.Reader and calls onData when a
// data element that matches the target path is encountered.
func MatchLoad(r io.Reader, onData func(path jp.Expr, data any), targets ...jp.Expr) error {
	_ = "STUB: not implemented"
	return nil
}
