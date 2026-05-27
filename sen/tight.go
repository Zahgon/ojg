// Copyright (c) 2021, Peter Ohler, All rights reserved.

package sen

import (
	"reflect"
)

func tightDefault(wr *Writer, data any, _ int) { _ = "STUB: not implemented"; return }

// Not much should get here except Map, Complex and un-decomposable
// values.

func tightArray(wr *Writer, n []any, _ int) { _ = "STUB: not implemented"; return }

func tightObject(wr *Writer, n map[string]any, _ int) { _ = "STUB: not implemented"; return }

func tightSortObject(wr *Writer, n map[string]any, _ int) { _ = "STUB: not implemented"; return }

func (wr *Writer) tightStruct(rv reflect.Value, si *sinfo) { _ = "STUB: not implemented"; return }

// Check for nil of any type

func (wr *Writer) tightSlice(rv reflect.Value, si *sinfo) { _ = "STUB: not implemented"; return }

func (wr *Writer) tightMap(rv reflect.Value, si *sinfo) { _ = "STUB: not implemented"; return }
