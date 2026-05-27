// Copyright (c) 2021, Peter Ohler, All rights reserved.

package alt

import (
	"time"
)

// TimeTolerance is the tolerance when comparing time elements
var TimeTolerance = time.Millisecond

// Path is a list of keys that can be either a string, int, or nil. Strings
// are used for keys in a map, ints are for indexes to a slice/array, and nil
// is a wildcard that matches either.
type Path []any

// String representation of the Path.
func (p Path) String() string { _ = "STUB: not implemented"; return "" }

// Diff returns the paths to the differences between two values. Any ignore
// paths are ignored in the comparison.
func Diff(v0, v1 any, ignores ...Path) (diffs []Path) { _ = "STUB: not implemented"; return nil }

// Compare returns a path to the first difference encountered between two
// values. Any ignore paths are ignored in the comparison.
func Compare(v0, v1 any, ignores ...Path) Path { _ = "STUB: not implemented"; return *new(Path) }

// Match returns true if all elements in the fingerprint match those in
// target. Fields in target but not in the fingerprint are ignored. An
// explicit nil in the fingerprint will match either a nil in the target or a
// missing value in the target.
func Match(fingerprint, target any) bool { _ = "STUB: not implemented"; return false }

func diff(v0, v1 any, one bool, ignores ...Path) (diffs []Path) {
	_ = "STUB: not implemented"
	return nil
}

// TBD optimize by a more direct compare of fields

func asInt(v any) (i int64, ok bool) { _ = "STUB: not implemented"; return 0, false }

func asFloat(v any) (f float64, ok bool) { _ = "STUB: not implemented"; return 0, false }

func ignoreIndex(i int, ignores []Path) bool { _ = "STUB: not implemented"; return false }

// wildcard, matches any index

func ignoreKey(k string, ignores []Path) bool { _ = "STUB: not implemented"; return false }

// wildcard, matches any index
