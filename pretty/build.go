// Copyright (c) 2021, Peter Ohler, All rights reserved.

package pretty

import (
	"time"

	"github.com/ohler55/ojg/gen"
)

func (w *Writer) build(data any) (n *node) { _ = "STUB: not implemented"; return nil }

// TBD OmitNil and OmitEmpty

// TBD OmitNil and OmitEmpty

func (w *Writer) buildNull() *node { _ = "STUB: not implemented"; return nil }

func (w *Writer) buildBool(v bool) (n *node) { _ = "STUB: not implemented"; return nil }

func (w *Writer) buildInt(v int64) (n *node) { _ = "STUB: not implemented"; return nil }

func (w *Writer) buildFloat32(v float32) (n *node) { _ = "STUB: not implemented"; return nil }

func (w *Writer) buildFloat64(v float64) (n *node) { _ = "STUB: not implemented"; return nil }

func (w *Writer) buildStringNode(v string) (n *node) { _ = "STUB: not implemented"; return nil }

func (w *Writer) buildTimeNode(v time.Time) (n *node) { _ = "STUB: not implemented"; return nil }

func (w *Writer) buildArrayNode(v []any) (n *node) { _ = "STUB: not implemented"; return nil }

// []

// space

// comma

func (w *Writer) buildGenArrayNode(v gen.Array) (n *node) { _ = "STUB: not implemented"; return nil }

// []

// space

// comma

func (w *Writer) buildMapNode(v map[string]any) (n *node) { _ = "STUB: not implemented"; return nil }

// {}

// build key

// space

// comma

// key, colon, space, value

func (w *Writer) buildGenMapNode(v gen.Object) (n *node) { _ = "STUB: not implemented"; return nil }

// {}

// build key

// space

// comma

// key, colon, space, value
