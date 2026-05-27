// Copyright (c) 2020, Peter Ohler, All rights reserved.

package oj

import (
	"io"
	"reflect"

	"github.com/ohler55/ojg"
)

const (
	spaces = "\n                                                                " +
		"                                                                "
	tabs = "\n\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t\t"
)

// Writer is a JSON writer that includes a reused buffer for reduced
// allocations for repeated encoding calls.
type Writer struct {
	ojg.Options
	buf           []byte
	w             io.Writer
	findex        byte
	strict        bool
	appendArray   func(wr *Writer, data []any, depth int)
	appendObject  func(wr *Writer, data map[string]any, depth int)
	appendDefault func(wr *Writer, data any, depth int)
	appendString  func(buf []byte, s string, htmlSafe bool) []byte
}

// JSON writes data, JSON encoded. On error, an empty string is returned.
func (wr *Writer) JSON(data any) string { _ = "STUB: not implemented"; return "" }

// MustJSON writes data, JSON encoded as a []byte and not a string like the
// JSON() function. On error a panic is called with the error. The returned
// buffer is the Writer buffer and is reused on the next call to write. If
// returned value is to be preserved past a second invocation then the buffer
// should be copied.
func (wr *Writer) MustJSON(data any) []byte { _ = "STUB: not implemented"; return nil }

// Write a JSON string for the data provided.
func (wr *Writer) Write(w io.Writer, data any) (err error) { _ = "STUB: not implemented"; return nil }

// MustWrite a JSON string for the data provided. If an error occurs panic is
// called with the error.
func (wr *Writer) MustWrite(w io.Writer, data any) { _ = "STUB: not implemented"; return }

func (wr *Writer) calcFieldsIndex() { _ = "STUB: not implemented"; return }

func (wr *Writer) appendJSON(data any, depth int) { _ = "STUB: not implemented"; return }

// go marshal treats a nil slice as a special case different from an
// empty slice. Seems kind of odd but here is the check.

func appendDefault(wr *Writer, data any, depth int) { _ = "STUB: not implemented"; return }

func appendArray(wr *Writer, n []any, depth int) { _ = "STUB: not implemented"; return }

func appendObject(wr *Writer, n map[string]any, depth int) { _ = "STUB: not implemented"; return }

func appendSortObject(wr *Writer, n map[string]any, depth int) { _ = "STUB: not implemented"; return }

func (wr *Writer) appendStruct(rv reflect.Value, depth int, si *sinfo) {
	_ = "STUB: not implemented"
	return
}

// Check for nil of any type

func (wr *Writer) appendSlice(rv reflect.Value, depth int, si *sinfo) {
	_ = "STUB: not implemented"
	return
}

func (wr *Writer) appendMap(rv reflect.Value, depth int, si *sinfo) {
	_ = "STUB: not implemented"
	return
}
