// Copyright (c) 2020, Peter Ohler, All rights reserved.

package gen

// Sort if true sorts Object keys on output.
var Sort = false

// Object is a map of Nodes with string keys.
type Object map[string]Node

// String returns a string representation of the Node.
func (n Object) String() string { _ = "STUB: not implemented"; return "" }

// Alter the Object into a simple map[string]any.
func (n Object) Alter() any { _ = "STUB: not implemented"; return *new(any) }

// Simplify creates a simplified version of the Node as a
// map[string]any.
func (n Object) Simplify() any { _ = "STUB: not implemented"; return *new(any) }

// Dup creates a deep duplicate of the Node.
func (n Object) Dup() Node { _ = "STUB: not implemented"; return *new(Node) }

// Empty returns true if the Object is empty.
func (n Object) Empty() bool { _ = "STUB: not implemented"; return false }
