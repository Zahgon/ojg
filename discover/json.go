// Copyright (c) 2025, Peter Ohler, All rights reserved.

package discover

import (
	"io"
)

// JSON finds occurrence of JSON documents that are either maps or arrays. The
// callback function should return true to stop discovering.
func JSON(buf []byte, cb func(value any) (stop bool)) { _ = "STUB: not implemented"; return }

// ReadJSON finds occurrence of JSON documents that are either maps or arrays in
// a stream. The callback function should return true to stop discovering.
func ReadJSON(r io.Reader, cb func(value any) bool) { _ = "STUB: not implemented"; return }
