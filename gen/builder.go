// Copyright (c) 2020, Peter Ohler, All rights reserved.

package gen

// Builder is assists in build a more complex Node.
type Builder struct {
	stack  []Node
	starts []int
}

// Reset clears the the Builder of previous built nodes.
func (b *Builder) Reset() { _ = "STUB: not implemented"; return }

// MustObject adds an object to the builder. A key is required if adding to a
// parent object.
func (b *Builder) MustObject(key ...string) { _ = "STUB: not implemented"; return }

// Object adds an object to the builder. A key is required if adding to a
// parent object.
func (b *Builder) Object(key ...string) error { _ = "STUB: not implemented"; return nil }

// MustArray adds an array to the builder. A key is required if adding to a
// parent object.
func (b *Builder) MustArray(key ...string) { _ = "STUB: not implemented"; return }

// Array adds an array to the builder. A key is required if adding to a parent
// object.
func (b *Builder) Array(key ...string) error { _ = "STUB: not implemented"; return nil }

// MustValue adds a Node to the builder. A key is required if adding to a
// parent object.
func (b *Builder) MustValue(value Node, key ...string) { _ = "STUB: not implemented"; return }

// Value adds a Node to the builder. A key is required if adding to a parent
// object.
func (b *Builder) Value(value Node, key ...string) error { _ = "STUB: not implemented"; return nil }

// Pop close a parent Object or Array Node.
func (b *Builder) Pop() { _ = "STUB: not implemented"; return }

// array

// PopAll close all parent Object or Array Nodes.
func (b *Builder) PopAll() { _ = "STUB: not implemented"; return }

// Result returns the current built Node.
func (b *Builder) Result() (result Node) { _ = "STUB: not implemented"; return *new(Node) }
