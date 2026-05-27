// Copyright (c) 2020, Peter Ohler, All rights reserved.

package oj

import (
	"io"
)

const stackMinSize = 32 // for container stack { or [

// Validator is a reusable JSON validator. It can be reused for multiple
// validations or parsings which allows buffer reuse for a performance
// advantage.
type Validator struct {
	tracker

	// This and the Parser use the same basic code but without the
	// building. It is a copy since adding the conditionals needed to avoid
	// building results add 15 to 20% overhead. An additional improvement could
	// be made by not tracking line and column but that would make it
	// validation much less useful.
	stack    []byte // { or [
	ri       int    // read index for null, false, and true
	mode     string
	nextMode string

	// OnlyOne returns an error if more than one JSON is in the string or
	// stream.
	OnlyOne bool
}

// Validate a JSON encoded byte slice.
func (p *Validator) Validate(buf []byte) (err error) { _ = "STUB: not implemented"; return nil }

// Skip BOM if present.

// ValidateReader a JSON stream. An error is returned if not valid JSON.
func (p *Validator) ValidateReader(r io.Reader) error { _ = "STUB: not implemented"; return nil }

// Skip BOM if present.

func (p *Validator) validateBuffer(buf []byte, last bool) error {
	_ = "STUB: not implemented"
	return nil
}

// nothing to do

// valid finishing maps are one byte longer
