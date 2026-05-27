// Copyright (c) 2025, Peter Ohler, All rights reserved.

package discover

import "io"

// Find potential occurrence of SEN documents that are either maps or
// arrays. This is a best effort search to find potential SEN documents. It is
// possible that document will not parse without errors. The callback function
// should return a true back return value to back up to the next open
// character after the current start. If back is false scanning continues
// after the end of the found section. If stop is true then no further
// scanning is attempted and the function returns.
func Find(buf []byte, cb func(found []byte) (back, stop bool)) { _ = "STUB: not implemented"; return }

func find(
	buf []byte,
	cb func(found []byte) (back, stop bool),
	more func(buf []byte, start, i int) ([]byte, int, int, bool)) {
	_ = "STUB: not implemented"
	return
}

// no change

// Read finds potential occurrence of SEN documents that are either maps or
// arrays in a stream. This is a best effort search to find potential SEN
// documents. It is possible that document will not parse without errors. The
// callback function should return a true back return value to back up to the
// next open character after the current start. If back is false scanning
// continues after the end of the found section. If stop is true then no
// further scanning is attempted and the function returns.
func Read(r io.Reader, cb func(b []byte) (back, stop bool)) { _ = "STUB: not implemented"; return }
