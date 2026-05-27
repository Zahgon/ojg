// Copyright (c) 2020, Peter Ohler, All rights reserved.

package tt

// ShortReader readons only the designated amount and then returns an
// error.
type ShortReader struct {
	Max     int
	Content []byte
	pos     int
}

// Read the next batch of bytes.
func (r *ShortReader) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }
