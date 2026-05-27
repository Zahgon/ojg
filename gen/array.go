// Copyright (c) 2020, Peter Ohler, All rights reserved.

package gen

// Array represents an array of nodes.
type Array []Node

// EmptyArray is a array of nodes of zero length.
var EmptyArray = Array{}

func (n Array) String() string { _ = "STUB: not implemented"; return "" }

// Alter the array into a simple []any.
func (n Array) Alter() any { _ = "STUB: not implemented"; return *new(any) }

// Simplify creates a simplified version of the Node as a []any.
func (n Array) Simplify() any { _ = "STUB: not implemented"; return *new(any) }

// Dup creates a deep duplicate of the Node.
func (n Array) Dup() Node { _ = "STUB: not implemented"; return *new(Node) }

// Empty returns true if the Array is empty.
func (n Array) Empty() bool { _ = "STUB: not implemented"; return false }
