// Copyright (c) 2020, Peter Ohler, All rights reserved.

package alt

var emptySlice = []any{}

// Builder is a basic type builder. It uses a stack model to build where maps
// (objects) and slices (arrays) add pushed on the stack and closed with a
// pop.
type Builder struct {
	stack  []any
	starts []int
}

// Reset the builder.
func (b *Builder) Reset() { _ = "STUB: not implemented"; return }

// Object pushs a map[string]any onto the stack. A key must be
// provided if the top of the stack is an object (map) and must not be
// provided if the op of the stack is an array or slice.
func (b *Builder) Object(key ...string) error { _ = "STUB: not implemented"; return nil }

// Array pushs a []any onto the stack. A key must be provided if the
// top of the stack is an object (map) and must not be provided if the op of
// the stack is an array or slice.
func (b *Builder) Array(key ...string) error { _ = "STUB: not implemented"; return nil }

// Value pushs a value onto the stack. A key must be provided if the top of
// the stack is an object (map) and must not be provided if the op of the
// stack is an array or slice.
func (b *Builder) Value(value any, key ...string) error { _ = "STUB: not implemented"; return nil }

// Pop the stack, closing an array or object.
func (b *Builder) Pop() { _ = "STUB: not implemented"; return }

// array

// PopAll repeats Pop until all open arrays or objects are closed.
func (b *Builder) PopAll() { _ = "STUB: not implemented"; return }

// Result of the builder is returned. This is the first item pushed on to the
// stack.
func (b *Builder) Result() (result any) { _ = "STUB: not implemented"; return *new(any) }
