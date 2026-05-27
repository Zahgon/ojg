// Copyright (c) 2021, Peter Ohler, All rights reserved.

package sen

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

// Writer is a SEN writer that includes a reused buffer for reduced
// allocations for repeated encoding calls.
type Writer struct {
	ojg.Options
	buf           []byte
	w             io.Writer
	appendArray   func(wr *Writer, data []any, depth int)
	appendObject  func(wr *Writer, data map[string]any, depth int)
	appendDefault func(wr *Writer, data any, depth int)
	appendString  func(buf []byte, s string, htmlSafe bool) []byte
	findex        byte
	needSep       bool
}

// SEN writes data, SEN encoded. On error, an empty string is returned.
func (wr *Writer) SEN(data any) string { _ = "STUB: not implemented"; return "" }

// MustSEN writes data, SEN encoded as a []byte and not a string like the
// SEN() function. On error a panic is called with the error. The returned
// buffer is the Writer buffer and is reused on the next call to write. If
// returned value is to be preserved past a second invocation then the buffer
// should be copied.
func (wr *Writer) MustSEN(data any) []byte { _ = "STUB: not implemented"; return nil }

// Write a SEN string for the data provided.
func (wr *Writer) Write(w io.Writer, data any) (err error) { _ = "STUB: not implemented"; return nil }

// MustWrite a SEN string for the data provided. If an error occurs panic is
// called with the error.
func (wr *Writer) MustWrite(w io.Writer, data any) { _ = "STUB: not implemented"; return }

func (wr *Writer) calcFieldsIndex() { _ = "STUB: not implemented"; return }

func (wr *Writer) appendSEN(data any, depth int) { _ = "STUB: not implemented"; return }

func appendDefault(wr *Writer, data any, depth int) { _ = "STUB: not implemented"; return }

// Not much should get here except Complex and non-decomposable
// values.

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
