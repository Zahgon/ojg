// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

var fnMap = map[string]Fn{}

// Fn encapsulates the information about a formula function in the package.
type Fn struct {
	Name     string
	Eval     func(root map[string]any, at any, args ...any) any
	Args     []any
	Desc     string
	Compile  func(*Fn)
	compiled bool
}

// Define a function for assembly use.
func Define(f *Fn) { _ = "STUB: not implemented"; return }

// FnDocs returns the documentation for all function.
func FnDocs() map[string]string { _ = "STUB: not implemented"; return nil }

// NewFn create a new named function of the named behavior.
func NewFn(name string) (fn *Fn) { _ = "STUB: not implemented"; return nil }

// Simplify a function in to simple types that can be encodes as JSON or SEN.
func (f *Fn) Simplify() any { _ = "STUB: not implemented"; return *new(any) }

// String return a string representation of the function.
func (f *Fn) String() string { _ = "STUB: not implemented"; return "" }

func (f *Fn) compile() { _ = "STUB: not implemented"; return }

func evalArg(root map[string]any, at, arg any) (val any) {
	_ = "STUB: not implemented"
	return *new(any)
}
