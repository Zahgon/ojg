// Copyright (c) 2020, Peter Ohler, All rights reserved.

package oj

import (
	"reflect"
)

func tightDefault(wr *Writer, data any, _ int) { _ = "STUB: not implemented"; return }

func tightArray(wr *Writer, n []any, _ int) { _ = "STUB: not implemented"; return }

func tightObject(wr *Writer, n map[string]any, _ int) { _ = "STUB: not implemented"; return }

func tightSortObject(wr *Writer, n map[string]any, _ int) { _ = "STUB: not implemented"; return }

func (wr *Writer) tightStruct(rv reflect.Value, si *sinfo) { _ = "STUB: not implemented"; return }

// Check for nil of any type

func (wr *Writer) tightSlice(rv reflect.Value, si *sinfo) { _ = "STUB: not implemented"; return }

func (wr *Writer) tightMap(rv reflect.Value, si *sinfo) { _ = "STUB: not implemented"; return }
