// Copyright (c) 2020, Peter Ohler, All rights reserved.

package gen

import (
	"strings"
	"time"
)

// TimeFormat defines how time is encoded. Options are to use a time. layout
// string format such as time.RFC3339Nano, "second" for a decimal
// representation, "nano" for a an integer.
var TimeFormat = ""

// TimeWrap if not empty encoded time as an object with a single member. For
// example if set to "@" then and TimeFormat is RFC3339Nano then the encoded
// time will look like '{"@":"2020-04-12T16:34:04.123456789Z"}'
var TimeWrap = ""

// Time is a time.Time Node.
type Time time.Time

// String returns a string representation of the Node.
func (n Time) String() string { _ = "STUB: not implemented"; return "" }

// Alter returns the backing time.Time value of the Node.
func (n Time) Alter() any {
	_ = "STUB: not implemented"
	return *

	// Simplify returns the backing time.Time value of the Node.
	new(any)
}

func (n Time) Simplify() any {
	_ = "STUB: not implemented"
	return *

	// Dup returns the backing time.Time value of the Node.
	new(any)
}

func (n Time) Dup() Node {
	_ = "STUB: not implemented"

	// Empty returns false.
	return *new(Node)
}

func (n Time) Empty() bool { _ = "STUB: not implemented"; return false }

func (n Time) buildString(b *strings.Builder) { _ = "STUB: not implemented"; return }

// Decimal format but float is not accurate enough so build the output
// in two parts.
