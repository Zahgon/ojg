// Copyright (c) 2020, Peter Ohler, All rights reserved.

package jp

// X creates an empty Expr.
func X() Expr {
	_ = "STUB: not implemented"

	// A creates an Expr with a At (@) fragment.
	return *new(Expr)
}

func A() Expr {
	_ = "STUB: not implemented"
	return *

	// B creates an Expr with a Bracket fragment.
	new(Expr)
}

func B() Expr {
	_ = "STUB: not implemented"
	return *

	// C creates an Expr with a Child fragment.
	new(Expr)
}

func C(key string) Expr {
	_ = "STUB: not implemented"
	return *

	// D creates an Expr with a recursive Descent fragment.
	new(Expr)
}

func D() Expr {
	_ = "STUB: not implemented"
	return *

	// F creates an Expr with a Filter fragment.
	new(Expr)
}

func F(e *Equation) Expr {
	_ = "STUB: not implemented"
	return *

	// N creates an Expr with an Nth fragment.
	new(Expr)
}

func N(n int) Expr {
	_ = "STUB: not implemented"
	return *

	// R creates an Expr with a Root fragment.
	new(Expr)
}

func R() Expr {
	_ = "STUB: not implemented"
	return *

	// S creates an Expr with a Slice fragment.
	new(Expr)
}

func S(start int, rest ...int) Expr { _ = "STUB: not implemented"; return *new(Expr) }

// U creates an Expr with an Union fragment.
func U(keys ...any) Expr { _ = "STUB: not implemented"; return *new(Expr) }

// W creates an Expr with a Wildcard fragment.
func W() Expr { _ = "STUB: not implemented"; return *new(Expr) }

// A appends an At fragment to the Expr.
func (x Expr) A() Expr {
	_ = "STUB: not implemented"
	return *

	// At appends an At fragment to the Expr.
	new(Expr)
}

func (x Expr) At() Expr {
	_ = "STUB: not implemented"
	return *

	// B appends a Bracket fragment to the Expr.
	new(Expr)
}

func (x Expr) B() Expr { _ = "STUB: not implemented"; return *new(Expr) }

// C appends a Child fragment to the Expr.
func (x Expr) C(key string) Expr { _ = "STUB: not implemented"; return *new(Expr) }

// Child appends a Child fragment to the Expr.
func (x Expr) Child(key string) Expr { _ = "STUB: not implemented"; return *new(Expr) }

// D appends a recursive Descent fragment to the Expr.
func (x Expr) D() Expr { _ = "STUB: not implemented"; return *new(Expr) }

// Descent appends a recursive Descent fragment to the Expr.
func (x Expr) Descent() Expr { _ = "STUB: not implemented"; return *new(Expr) }

// F appends a Filter fragment to the Expr.
func (x Expr) F(e *Equation) Expr { _ = "STUB: not implemented"; return *new(Expr) }

// Filter appends a Filter fragment to the Expr.
func (x Expr) Filter(e *Equation) Expr { _ = "STUB: not implemented"; return *new(Expr) }

// N appends an Nth fragment to the Expr.
func (x Expr) N(n int) Expr {
	_ = "STUB: not implemented"
	return *

	// Nth appends an Nth fragment to the Expr.
	new(Expr)
}

func (x Expr) Nth(n int) Expr {
	_ = "STUB: not implemented"
	return *

	// R appends a Root fragment to the Expr.
	new(Expr)
}

func (x Expr) R() Expr { _ = "STUB: not implemented"; return *new(Expr) }

// Root appends a Root fragment to the Expr.
func (x Expr) Root() Expr { _ = "STUB: not implemented"; return *new(Expr) }

// S appends a Slice fragment to the Expr.
func (x Expr) S(start int, rest ...int) Expr { _ = "STUB: not implemented"; return *new(Expr) }

// Slice appends a Slice fragment to the Expr.
func (x Expr) Slice(start int, rest ...int) Expr { _ = "STUB: not implemented"; return *new(Expr) }

// U appends a Union fragment to the Expr.
func (x Expr) U(keys ...any) Expr { _ = "STUB: not implemented"; return *new(Expr) }

// Union appends a Union fragment to the Expr.
func (x Expr) Union(keys ...any) Expr { _ = "STUB: not implemented"; return *new(Expr) }

// W appends a Wildcard fragment to the Expr.
func (x Expr) W() Expr { _ = "STUB: not implemented"; return *new(Expr) }

// Wildcard appends a Wildcard fragment to the Expr.
func (x Expr) Wildcard() Expr { _ = "STUB: not implemented"; return *new(Expr) }
