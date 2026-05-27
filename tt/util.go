// Copyright (c) 2020, Peter Ohler, All rights reserved.

package tt

import (
	"strings"
	"testing"
)

type call struct {
	fn   string
	file string
	line int
}

func finishFail(t *testing.T, b *strings.Builder, args []any) { _ = "STUB: not implemented"; return }

func stackFill(b *strings.Builder) { _ = "STUB: not implemented"; return }

func isNil(v any) bool { _ = "STUB: not implemented"; return false }

func asInt(v any) (i int64, ok bool) { _ = "STUB: not implemented"; return 0, false }

func asFloat(v any) (f float64, ok bool) { _ = "STUB: not implemented"; return 0, false }

func asString(v any) (s string, ok bool) { _ = "STUB: not implemented"; return "", false }
