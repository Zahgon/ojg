// Copyright (c) 2020, Peter Ohler, All rights reserved.

package tt

import (
	"testing"
)

// Equal return true if two values are equal and fails a test if not equal.
func Equal(t *testing.T, expect, actual any, args ...any) (eq bool) {
	_ = "STUB: not implemented"
	return false
}

// NotEqual return true if two values are not equal and fails a test if equal.
func NotEqual(t *testing.T, expect, actual any, args ...any) (eq bool) {
	_ = "STUB: not implemented"
	return false
}

func valuesEqual(expect, actual any) (eq bool) { _ = "STUB: not implemented"; return false }

/*
	if !eq {
			if !eq {
				tx, ta = colorizeStrings(tx, ta)
				expect = tx
				actual = ta
			}
	}
*/
