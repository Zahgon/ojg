// Copyright (c) 2020, Peter Ohler, All rights reserved.

package jp

type nothing int

const userOpCode = 'U'

var (
	// Lower precedence is evaluated first.
	eq     = &op{prec: 3, code: '=', name: "==", cnt: 2}
	neq    = &op{prec: 3, code: 'n', name: "!=", cnt: 2}
	lt     = &op{prec: 3, code: '<', name: "<", cnt: 2}
	gt     = &op{prec: 3, code: '>', name: ">", cnt: 2}
	lte    = &op{prec: 3, code: 'l', name: "<=", cnt: 2}
	gte    = &op{prec: 3, code: 'g', name: ">=", cnt: 2}
	or     = &op{prec: 4, code: '|', name: "||", cnt: 2}
	and    = &op{prec: 4, code: '&', name: "&&", cnt: 2}
	not    = &op{prec: 0, code: '!', name: "!", cnt: 1}
	add    = &op{prec: 2, code: '+', name: "+", cnt: 2}
	sub    = &op{prec: 2, code: '-', name: "-", cnt: 2}
	mult   = &op{prec: 1, code: '*', name: "*", cnt: 2}
	divide = &op{prec: 1, code: '/', name: "/", cnt: 2}
	get    = &op{prec: 0, code: 'G', name: "get", cnt: 1}
	in     = &op{prec: 3, code: 'i', name: "in", cnt: 2}
	empty  = &op{prec: 3, code: 'e', name: "empty", cnt: 2}
	rx     = &op{prec: 3, code: '~', name: "~=", cnt: 2}
	rxa    = &op{prec: 3, code: '~', name: "=~", cnt: 2}
	has    = &op{prec: 3, code: 'h', name: "has", cnt: 2}
	exists = &op{prec: 3, code: 'x', name: "exists", cnt: 2}
	// functions
	length = &op{prec: 0, code: 'L', name: "length", cnt: 1}
	count  = &op{prec: 0, code: 'C', name: "count", cnt: 1, getLeft: true}
	match  = &op{prec: 0, code: 'M', name: "match", cnt: 2}
	search = &op{prec: 0, code: 'S', name: "search", cnt: 2}

	// group is for an equation inside () so it represents the (). It should
	// not be in the opMap.
	group = &op{prec: 0, code: '(', name: "(", cnt: 1}

	opMap = map[string]*op{
		eq.name:     eq,
		neq.name:    neq,
		lt.name:     lt,
		gt.name:     gt,
		lte.name:    lte,
		gte.name:    gte,
		or.name:     or,
		and.name:    and,
		not.name:    not,
		add.name:    add,
		sub.name:    sub,
		mult.name:   mult,
		divide.name: divide,
		in.name:     in,
		empty.name:  empty,
		has.name:    has,
		exists.name: exists,
		rx.name:     rx,
		rxa.name:    rx,

		length.name: length,
		count.name:  count,
		match.name:  match,
		search.name: search,
	}
	// Nothing can be used in scripts to indicate no value as in a script such
	// as [?(@.x == Nothing)] this indicates there was no value as @.x. It is
	// the same as [?(@.x has false)] or [?(@.x exists false)].
	Nothing = nothing(0)
)

type op struct {
	name     string
	uniFun   func(arg any) any
	duoFun   func(left, right any) any
	prec     byte
	cnt      byte
	code     byte
	getLeft  bool
	getRight bool
}

type precBuf struct {
	prec byte
	buf  []byte
}

type multivalue []any

type got struct {
	value any
}

// Script represents JSON Path script used in filters as well.
type Script struct {
	template []any
}

// NewScript parses the string argument and returns a script or an error.
func NewScript(str string) (s *Script, err error) { _ = "STUB: not implemented"; return nil, nil }

// MustNewScript parses the string argument and returns a script or an error.
func MustNewScript(str string) (s *Script) { _ = "STUB: not implemented"; return nil }

// Append a string representation of the fragment to the buffer and then
// return the expanded buffer.
func (s *Script) Append(buf []byte) []byte { _ = "STUB: not implemented"; return nil }

// String representation of the script.
func (s *Script) String() string { _ = "STUB: not implemented"; return "" }

// Match returns true if the script returns true when evaluated against the
// data argument.
func (s *Script) Match(data any) bool { _ = "STUB: not implemented"; return false }

// Eval is primarily used by the Expr parser but is public for testing.
func (s *Script) Eval(stack, data any) any { _ = "STUB: not implemented"; return *new(any) }

func (s *Script) evalWithRoot(stack, data, root any) (any, Expr) {
	_ = "STUB: not implemented"
	return *new(any), *new(Expr)
}

// Eval script for each member of the list.

// Check for functions like 'count'.

// TBD one more for getRight once function extensions are supported

// The most common pattern is [?(@.child == value)] where
// the operation and value vary but the @.child is the
// most widely used. For that reason an optimization is
// included for that condition of a one level child lookup
// path.

func normalize(v any) any {
	_ = "STUB: not implemented"
	// handle special values
	return *new(any)
}

// optimize for already normalized values

// handle inter-convertible values

// recursively handle pointers

// handle named types that implement common types

func expandStack(stack []any, mi int) []any { _ = "STUB: not implemented"; return nil }

func evalStack(sstack []any) []any { _ = "STUB: not implemented"; return nil }

// a value, not an op

// If one is a boolean true then true.

// If both are a boolean true then true else false.

// Inspect the script.
func (s *Script) Inspect() *Form { _ = "STUB: not implemented"; return nil }

func nextForm(st []any) (any, []any) { _ = "STUB: not implemented"; return *new(any), nil }

func (s *Script) appendOp(o *op, left, right any) (pb *precBuf) {
	_ = "STUB: not implemented"
	return nil
}

func (s *Script) appendValue(buf []byte, v any, prec byte) []byte {
	_ = "STUB: not implemented"
	return nil
}

var builtInNames = map[string]bool{
	"==":     true,
	"!=":     true,
	"<":      true,
	">":      true,
	"<=":     true,
	">=":     true,
	"||":     true,
	"&&":     true,
	"!":      true,
	"+":      true,
	"-":      true,
	"*":      true,
	"/":      true,
	"get":    true,
	"in":     true,
	"empty":  true,
	"~=":     true,
	"=~":     true,
	"has":    true,
	"exists": true,
	"length": true,
	"count":  true,
	"match":  true,
	"search": true,
	"true":   true,
	"false":  true,
	"null":   true,
}

// RegisterUnaryFunction registers a unary function for scripts. The 'get'
// argument if true indicates a get operation to provide the argument to the
// provided function otherwise the first match is used. Names must be alpha
// characters only.
func RegisterUnaryFunction(name string, get bool, f func(arg any) any) {
	_ = "STUB: not implemented"
	return
}

// RegisterBinaryFunction registers a function that takes two argument for
// scripts. The 'getLeft' and 'getRight' arguments if true indicates a get
// operation to provide the argument to the provided function otherwise the
// first match is used. Names must be alpha characters only.
func RegisterBinaryFunction(name string, getLeft, getRight bool, f func(left, right any) any) {
	_ = "STUB: not implemented"
	return
}
