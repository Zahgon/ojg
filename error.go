// Copyright (c) 2021, Peter Ohler, All rights reserved.

package ojg

// ErrorWithStack if true the Error() call will include the stack.
var ErrorWithStack = false

// Error struct to hold an error message and a stack trace.
type Error struct {
	msg   string
	stack []byte
}

// NewError creates a new Error instance, capturing the stack when created.
func NewError(r any) *Error { _ = "STUB: not implemented"; return nil }

// Error returns a string representation of the instance.
func (err *Error) Error() string { _ = "STUB: not implemented"; return "" }

// Stack returns the stack.
func (err *Error) Stack() []byte { _ = "STUB: not implemented"; return nil }
