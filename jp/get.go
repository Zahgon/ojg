// Copyright (c) 2020, Peter Ohler, All rights reserved.

package jp

import (
	"reflect"
)

const (
	fragIndexMask    = 0x0000ffff
	descentFlag      = 0x00010000
	descentChildFlag = 0x00020000

	// The standard math package fails to compile on 32bit architectures (ARM)
	// with an int overflow. Most likley due to math.MaxInt64 being defined as
	// 1<<63 - 1 which default to integer values. Since arrays are not likely
	// to be over 2147483647 on a 32 bit system that is set as the max end
	// specifier for a array range.
	maxEnd = 2147483647
)

type fragIndex int

// The easy way to implement the Get is to have each fragment handle the
// getting using recursion. The overhead of a go function call is rather high
// though so instead a pseudo call stack is implemented here that grows and
// shrinks as the getting takes place. The fragment index if placed on the
// stack as well mostly for a small degree of simplicity in what a few people
// might find a complex approach to the solution. Its at least twice as fast
// as the recursive function call approach and in some cases such as the
// recursive descent more than an order of magnitude faster.

// Get the elements of the data identified by the path.
func (x Expr) Get(data any) (results []any) { _ = "STUB: not implemented"; return nil }

// frag index

// must have at least a data element and a fragment index

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// first pass expands, second continues evaluation

// Put prev back and slide fi.

// last one

// Put prev back and slide fi.

// last one

// Put prev back and slide fi.

// last one

// TBD ??

// Put prev back and slide fi.

// last one

// Put prev back and slide fi.

// last one

// Put prev back and slide fi.

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// Free up anything still on the stack.

// First element of the data identified by the path.
func (x Expr) First(data any) any { _ = "STUB: not implemented"; return *new(any) }

// FirstFound element of the data identified by the path.
func (x Expr) FirstFound(data any) (any, bool) { _ = "STUB: not implemented"; return *new(any), false }

// frag index

// must have at least a data element and a fragment index

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// first pass expands, second continues evaluation

// Put prev back and slide fi.

// last one

// Put prev back and slide fi.

// last one

// Put prev back and slide fi.

// last one

// last one

// Put prev back and slide fi.

// last one

// Put prev back and slide fi.

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

// last one

func reflectGetChild(data any, key string) (v any, has bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func reflectGetFieldByKey(structValue reflect.Value, key string) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func reflectGetStructFieldByNameOrJsonTag(structValue reflect.Value, key string) (result reflect.StructField, ok bool) {
	_ = "STUB: not implemented"

	// match by field name or by json tag
	return *new(reflect.StructField), false
}

// -----------------------------------------------------------------
// The algorithm is breadth first search, one depth level at a time.
// Based on the original: ((Struct) reflect.Value).FieldByNameFunc()
// -----------------------------------------------------------------

// The 'current' and 'next' slices are work queues:
// 'current' lists the fields to visit on this depth level,
// and 'next' lists the fields on the next lower level.

// 'nextCount' records the number of times an embedded struct type has been
// encountered and considered for queueing in the 'next' slice.
// We only queue the first one, but we increment the count on each.
// If a struct type T can be reached more than once at a given depth level,
// then it annihilates itself and need not be considered at all when we
// process that next depth level.

// 'visited' records the structs that have been considered already.
// Note that embedded pointer fields can create cycles in the graph of
// reachable embedded types; 'visited' avoids following those cycles.
// It also avoids duplicated effort: if we didn't find the field in an
// embedded type T at level 2, we won't find it in one at level 4 either.

// Process all the fields at this depth, now listed in 'current'.
// The loop queues embedded fields found in 'next', for processing during the next
// iteration. The multiplicity of the 'current' field counts is recorded
// in 'count'; the multiplicity of the 'next' field counts is recorded in 'nextCount'.

// We've looked through this type before, at a higher level.
// That higher level would shadow the lower level we're now at,
// so this one can't be useful to us. Ignore it.

// Embedded field of type T or *T.

// Does it match?

// Potential match

// Name appeared multiple times at this level: annihilate.

// Queue embedded struct fields for processing with next level,
// but only if we haven't seen a match yet at this level and only
// if the embedded types haven't already been queued.

// here we are sure that the nested type is indeed a Struct

// exact multiple doesn't matter

// exact multiple doesn't matter

func reflectGetNth(data any, i int) (v any, has bool) {
	_ = "STUB: not implemented"
	return *new(any), false
}

func reflectGetWild(data any) (va []any) { _ = "STUB: not implemented"; return nil }

// Iterate in reverse order as that puts values on the stack in reverse.

func reflectGetWildOne(data any) (any, bool) { _ = "STUB: not implemented"; return *new(any), false }

func reflectGetSlice(data any, start, end, step int) (va []any) {
	_ = "STUB: not implemented"
	return nil
}
