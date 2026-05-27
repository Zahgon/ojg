// Copyright (c) 2024, Peter Ohler, All rights reserved.

package jp

// CompileScript if non-nil should return object that implments the Procedure
// interface. This function is called when a script notation bracketed by [(
// and )] is encountered. Note the string code argument will included the open
// and close parenthesis but not the square brackets.
var CompileScript func(code []byte) Procedure

// Proc is a script used as a procedure which is a script not limited to being
// a selector. While both Locate() and Walk() are supported the results may
// not be as expected since the procedure can modify the original
// data. Remove() is not supported with this fragment type.
type Proc struct {
	Procedure Procedure
	Script    []byte
}

// MustNewProc creates a new Proc and panics on error.
func MustNewProc(code []byte) (p *Proc) { _ = "STUB: not implemented"; return nil }

// String representation of the proc.
func (p *Proc) String() string { _ = "STUB: not implemented"; return "" }

// Append a fragment string representation of the fragment to the buffer
// then returning the expanded buffer.
func (p *Proc) Append(buf []byte, _, _ bool) []byte { _ = "STUB: not implemented"; return nil }

func (p *Proc) locate(pp Expr, data any, rest Expr, max int) (locs []Expr) {
	_ = "STUB: not implemented"
	return nil
}

// last one

// place holder

// Walk each element returned from the procedure call. Note that this may or
// may not correspond to the original data as the procedure can modify not only
// the elements in the original data but also the contents of each.
func (p *Proc) Walk(rest, path Expr, nodes []any, cb func(path Expr, nodes []any)) {
	_ = "STUB: not implemented"
	return
}
