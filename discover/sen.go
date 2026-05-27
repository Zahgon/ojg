// Copyright (c) 2025, Peter Ohler, All rights reserved.

package discover

import (
	"io"
)

// SEN finds occurrence of SEN documents that are either maps or arrays. The
// callback function should return true to stop discovering.
func SEN(buf []byte, cb func(value any) (stop bool)) { _ = "STUB: not implemented"; return }

// ReadSEN finds occurrence of SEN documents that are either maps or arrays in
// a stream. The callback function should return true to stop discovering.
func ReadSEN(r io.Reader, cb func(value any) bool) { _ = "STUB: not implemented"; return }
