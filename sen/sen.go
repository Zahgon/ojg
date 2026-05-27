// Copyright (c) 2021, Peter Ohler, All rights reserved.

// Package sen contains the SEN parsers and writers.
package sen

import (
	"io"
	"sync"

	"github.com/ohler55/ojg"
	"github.com/ohler55/ojg/alt"
	"github.com/ohler55/ojg/jp"
)

// Options is an alias for ojg.Options
type Options = ojg.Options

var (
	// DefaultOptions are the default options for the this package.
	DefaultOptions = ojg.DefaultOptions
	// BrightOptions are the bright color options.
	BrightOptions = ojg.BrightOptions
	// HTMLOptions are the options that can be used to encode as HTML JSON.
	HTMLOptions = ojg.HTMLOptions

	writerPool = sync.Pool{
		New: func() any {
			return &Writer{Options: DefaultOptions, buf: make([]byte, 0, 1024)}
		},
	}
	parserPool = sync.Pool{
		New: func() any {
			return &Parser{}
		},
	}
)

// Parse SEN into a simple type. Arguments are optional and can be a
// func(any) bool for callbacks or a chan any for chan based
// result delivery. The SEN parser will also Parse JSON.
//
// A func argument is the callback for the parser if processing multiple
// SENs. If no callback function is provided the processing is limited to
// only one SEN.
//
// A chan argument will be used to deliver parse results.
func Parse(buf []byte, args ...any) (any, error) { _ = "STUB: not implemented"; return *new(any), nil }

// MustParse SEN into a simple type. Arguments are optional and can be a
// func(any) bool for callbacks or a chan any for chan based
// result delivery. The SEN parser will also Parse JSON. Panics on error.
//
// A func argument is the callback for the parser if processing multiple
// SENs. If no callback function is provided the processing is limited to
// only one SEN.
//
// A chan argument will be used to deliver parse results.
func MustParse(buf []byte, args ...any) any { _ = "STUB: not implemented"; return *new(any) }

// ParseReader reads and parses SEN into a simple type. Arguments are optional
// and can be a func(any) bool for callbacks or a chan any for
// chan based result delivery. The SEN parser will also Parse JSON.
//
// A func argument is the callback for the parser if processing multiple
// SENs. If no callback function is provided the processing is limited to
// only one SEN.
//
// A chan argument will be used to deliver parse results.
func ParseReader(r io.Reader, args ...any) (data any, err error) {
	_ = "STUB: not implemented"
	return *new(any), nil
}

// MustParseReader reads and parses SEN into a simple type. Arguments are
// optional and can be a func(any) bool for callbacks or a chan
// any for chan based result delivery. The SEN parser will also Parse
// JSON. Panics on error.
//
// A func argument is the callback for the parser if processing multiple
// SENs. If no callback function is provided the processing is limited to
// only one SEN.
//
// A chan argument will be used to deliver parse results.
func MustParseReader(r io.Reader, args ...any) (data any) {
	_ = "STUB: not implemented"
	return *new(any)
}

// Unmarshal parses the provided JSON and stores the result in the value
// pointed to by vp.
func Unmarshal(data []byte, vp any, recomposer ...*alt.Recomposer) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// String returns a SEN string for the data provided. The data can be a simple
// type of nil, bool, int, floats, time.Time, []any, or
// map[string]any or a Node type, The args, if supplied can be an int
// as an indent, *ojg.Options, or a *Writer.
func String(data any, args ...any) string { _ = "STUB: not implemented"; return "" }

// Bytes returns a SEN []byte for the data provided. The data can be a simple
// type of nil, bool, int, floats, time.Time, []any, or
// map[string]any or a Node type, The args, if supplied can be an int
// as an indent, *ojg.Options, or a *Writer. The returned buffer is the Writer
// buffer and is reused on the next call to write. If returned value is to be
// preserved past a second invocation then the buffer should be copied.
func Bytes(data any, args ...any) []byte { _ = "STUB: not implemented"; return nil }

// Write SEN for the data provided. The data can be a simple type of nil,
// bool, int, floats, time.Time, []any, or map[string]any or a
// Node type, The args, if supplied can be an int as an indent, *ojg.Options,
// or a *Writer.
func Write(w io.Writer, data any, args ...any) (err error) { _ = "STUB: not implemented"; return nil }

// MustWrite SEN for the data provided. The data can be a simple type of nil,
// bool, int, floats, time.Time, []any, or map[string]any or a
// Node type, The args, if supplied can be an int as an indent, *ojg.Options,
// or a *Writer. Panics on error.
func MustWrite(w io.Writer, data any, args ...any) { _ = "STUB: not implemented"; return }

func pickWriter(arg any) (wr *Writer) { _ = "STUB: not implemented"; return nil }

// Match parses a SEN document and calls onData when a data element that
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
