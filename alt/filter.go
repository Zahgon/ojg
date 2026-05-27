// Copyright (c) 2023, Peter Ohler, All rights reserved.

package alt

// Filter is a simple filter for matching against arbitrary date.
type Filter map[string]any

// NewFilter creates a new filter from the spec which should be a map where
// the keys are simple paths of keys delimited by the dot ('.') character. An
// example is "top.child.grandchild". The matching will either match the key
// when the data is traversed directly or in the case of a slice the elements
// of the slice are also traversed. Generally a Filter is created and reused
// as there is some overhead in creating the Filter. An alternate format is a
// nested set of maps.
func NewFilter(spec map[string]any) Filter { _ = "STUB: not implemented"; return *new(Filter) }

func (f Filter) add(spec map[string]any) { _ = "STUB: not implemented"; return }

// Match returns true if the target matches the Filter.
func (f Filter) Match(data any) bool { _ = "STUB: not implemented"; return false }

func match(target, data any) (same bool) { _ = "STUB: not implemented"; return false }

// Simplify returns a simplified representation of the Filter.
func (f Filter) Simplify() any { _ = "STUB: not implemented"; return *new(any) }
