// Copyright (c) 2023, Peter Ohler, All rights reserved.

package jp

// Locate the values described by the Expr and return a slice of normalized
// paths to those values in the data. The returned slice is limited to the max
// specified. A max of 0 or less indicates there is no maximum.
func (x Expr) Locate(data any, max int) (locs []Expr) { _ = "STUB: not implemented"; return nil }

func locateNthChildHas(pp Expr, f Frag, v any, rest Expr, max int) (locs []Expr) {
	_ = "STUB: not implemented"
	// last one
	return nil
}

func locateAppendFrag(locs []Expr, pp Expr, f Frag) []Expr { _ = "STUB: not implemented"; return nil }

func locateContinueFrag(locs []Expr, cp Expr, v any, rest Expr, max int) []Expr {
	_ = "STUB: not implemented"
	return nil
}
