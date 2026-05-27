// Copyright (c) 2024, Peter Ohler, All rights reserved.

package jp

// TargetRest is used by the MatchHandler to associate a Target and Rest of a
// match search.
type TargetRest struct {
	Target Expr
	// Rest is set when a Filter is included in the initializing target. Since
	// Filters can only be evaluated when there is data for the evaluation a
	// traget with a Filter is split with the pre-filter portion and the rest
	// starting with the filter.
	Rest Expr
}

// PathHandler is a TokenHandler compatible with both the oj.TokenHandler and
// the sen.TokenHandler. Fields are public to allow derived types to access
// those fields.
type MatchHandler struct {
	Targets []*TargetRest
	Path    Expr
	Stack   []any
	OnData  func(path Expr, data any)
}

// NewMatchHandler creates a new MatchHandler.
func NewMatchHandler(onData func(path Expr, data any), targets ...Expr) *MatchHandler {
	_ = "STUB: not implemented"
	return nil
}

// Null is called when a JSON null is encountered.
func (h *MatchHandler) Null() {
	_ = "STUB: not implemented"

	// Bool is called when a JSON true or false is encountered.
	return
}

func (h *MatchHandler) Bool(v bool) {
	_ = "STUB: not implemented"

	// Int is called when a JSON integer is encountered.
	return
}

func (h *MatchHandler) Int(v int64) {
	_ = "STUB: not implemented"

	// Float is called when a JSON decimal is encountered that fits into a
	// float64.
	return
}

func (h *MatchHandler) Float(v float64) {
	_ = "STUB: not implemented"

	// Number is called when a JSON number is encountered that does not fit
	// into an int64 or float64.
	return
}

func (h *MatchHandler) Number(num string) { _ = "STUB: not implemented"; return }

// String is called when a JSON string is encountered.
func (h *MatchHandler) String(v string) {
	_ = "STUB: not implemented"

	// ObjectStart is called when a JSON object start '{' is encountered.
	return
}

func (h *MatchHandler) ObjectStart() { _ = "STUB: not implemented"; return }

// ObjectEnd is called when a JSON object end '}' is encountered.
func (h *MatchHandler) ObjectEnd() {
	_ = "STUB: not implemented"

	// Key is called when a JSON object key is encountered.
	return
}

func (h *MatchHandler) Key(k string) { _ = "STUB: not implemented"; return }

// ArrayStart is called when a JSON array start '[' is encountered.
func (h *MatchHandler) ArrayStart() { _ = "STUB: not implemented"; return }

// ArrayEnd is called when a JSON array end ']' is encountered.
func (h *MatchHandler) ArrayEnd() {
	_ = "STUB: not implemented"

	// AddValue is called when a leaf value is encountered.
	return
}

func (h *MatchHandler) AddValue(v any) { _ = "STUB: not implemented"; return }

func (h *MatchHandler) objArrayStart(v any, frag Frag) { _ = "STUB: not implemented"; return }

func (h *MatchHandler) objArrayEnd() { _ = "STUB: not implemented"; return }

func (h *MatchHandler) incNth() { _ = "STUB: not implemented"; return }

func (h *MatchHandler) checkRest(v any) (any, Expr, bool) {
	_ = "STUB: not implemented"
	return *new(any), *new(Expr), false
}

func (h *MatchHandler) pathMatch(leaf bool) bool { _ = "STUB: not implemented"; return false }
