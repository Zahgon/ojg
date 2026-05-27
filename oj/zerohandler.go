// Copyright (c) 2021, Peter Ohler, All rights reserved.

package oj

// ZeroHandler is a TokenHandler whose functions do nothing. It is used as an
// embedded member for TokenHandlers that don't care about all of the
// TokenHandler functions.
type ZeroHandler struct {
}

// Null is called when a JSON null is encountered.
func (z *ZeroHandler) Null() {
	_ = "STUB: not implemented"

	// Bool is called when a JSON true or false is encountered.
	return
}

func (z *ZeroHandler) Bool(bool) {
	_ = "STUB: not implemented"

	// Int is called when a JSON integer is encountered.
	return
}

func (z *ZeroHandler) Int(int64) {
	_ = "STUB: not implemented"

	// Float is called when a JSON decimal is encountered that fits into a
	// float64.
	return
}

func (z *ZeroHandler) Float(float64) {
	_ = "STUB: not implemented"

	// Number is called when a JSON number is encountered that does not fit
	// into an int64 or float64.
	return
}

func (z *ZeroHandler) Number(string) {
	_ = "STUB: not implemented"

	// String is called when a JSON string is encountered.
	return
}

func (z *ZeroHandler) String(string) {
	_ = "STUB: not implemented"

	// ObjectStart is called when a JSON object start '{' is encountered.
	return
}

func (z *ZeroHandler) ObjectStart() {
	_ = "STUB: not implemented"

	// ObjectEnd is called when a JSON object end '}' is encountered.
	return
}

func (z *ZeroHandler) ObjectEnd() {
	_ = "STUB: not implemented"

	// Key is called when a JSON object key is encountered.
	return
}

func (z *ZeroHandler) Key(string) {
	_ = "STUB: not implemented"

	// ArrayStart is called when a JSON array start '[' is encountered.
	return
}

func (z *ZeroHandler) ArrayStart() {
	_ = "STUB: not implemented"

	// ArrayEnd is called when a JSON array end ']' is encountered.
	return
}

func (z *ZeroHandler) ArrayEnd() { _ = "STUB: not implemented"; return }
