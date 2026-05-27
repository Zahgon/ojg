// Copyright (c) 2021, Peter Ohler, All rights reserved.

package pretty

import (
	"io"

	"github.com/ohler55/ojg"
)

const (
	nullStr  = "null"
	trueStr  = "true"
	falseStr = "false"
	spaces   = "\n                                                                " +
		"                                                                "
)

// Writer writes data in either JSON or SEN format using setting to determine
// the output.
type Writer struct {
	ojg.Options

	// Width is the suggested maximum width. In some cases it may not be
	// possible to stay within the specified width.
	Width int

	// MaxDepth is the maximum depth of an element on a single line.
	MaxDepth int

	// Align if true attempts to align elements of children in list.
	Align bool

	// SEN format if true otherwise JSON encoding.
	SEN bool

	buf []byte
	w   io.Writer
}

// Encode data. Any panics during encoding will cause an empty return but will
// not fail.The returned buffer is the Writer buffer and is reused on the next
// call to write. If returned value is to be preserved past a second
// invocation then the buffer should be copied.
func (w *Writer) Encode(data any) []byte { _ = "STUB: not implemented"; return nil }

// Marshal data. The same as Encode but a panics during encoding will result
// in an error return.
func (w *Writer) Marshal(data any) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// Write encoded data to the op.Writer. The returned buffer is the Writer
// buffer and is reused on the next call to write. If returned value is to be
// preserved past a second invocation then the buffer should be copied.
func (w *Writer) Write(wr io.Writer, data any) (err error) { _ = "STUB: not implemented"; return nil }

func (w *Writer) config(args []any) { _ = "STUB: not implemented"; return }

// use the default

func (w *Writer) encode(data any) (out []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *Writer) fill(n *node, depth int, flat bool) { _ = "STUB: not implemented"; return }

// Return true if not filled.
func (w *Writer) checkAlign(n *node, start int, comma, cs []byte) bool {
	_ = "STUB: not implemented"
	return false
}

func (w *Writer) alignArray(n *node, t *table, comma, cs []byte) { _ = "STUB: not implemented"; return }

func (w *Writer) alignMap(n *node, t *table, comma, cs []byte) { _ = "STUB: not implemented"; return }
