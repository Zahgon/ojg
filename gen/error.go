// Copyright (c) 2020, Peter Ohler, All rights reserved.

package gen

// ParseError represents a parse error.
type ParseError struct {
	Message string
	Line    int
	Column  int
}

// Error returns a string representation of the error.
func (err *ParseError) Error() string { _ = "STUB: not implemented"; return "" }
