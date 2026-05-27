// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

// Plan is an assembly plan that can be described by a JSON document or a SEN
// document. The format is much like LISP but with brackets instead of
// parenthesis. A plan is evaluated by evaluating the plan function which is
// usually an 'asm' function. The plan operates on a data map which is the
// root during evaluation. The source data is in the $.src and the expected
// assembled output should be in $.asm.
type Plan struct {
	Fn
}

// NewPlan creates new place from a simplified (JSON) encoding of the
// instance.
func NewPlan(plan []any) *Plan { _ = "STUB: not implemented"; return nil }

// Execute a plan.
func (p *Plan) Execute(root map[string]any) (err error) { _ = "STUB: not implemented"; return nil }
