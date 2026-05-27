// Copyright (c) 2021, Peter Ohler, All rights reserved.

package sen

import (
	"reflect"
	"sync"
)

const (
	maskByTag  = byte(0x01)
	maskExact  = byte(0x02) // exact key vs lowwer case first letter
	maskNested = byte(0x04)
	maskPretty = byte(0x08)
	maskMax    = byte(0x10)
)

type sinfo struct {
	rt     reflect.Type
	fields [16][]*finfo
}

var (
	structMut sync.Mutex
	// Keyed by the pointer to the type.
	structMap      = map[uintptr]*sinfo{}
	structEmptyMap = map[uintptr]*sinfo{}
)

// Non-locking version used in field creation.
func getTypeStruct(rt reflect.Type, embedded, omitEmpty bool) (st *sinfo) {
	_ = "STUB: not implemented"
	return nil
}

func getSinfo(v any, omitEmpty bool) (st *sinfo) { _ = "STUB: not implemented"; return nil }

func buildStruct(rt reflect.Type, x uintptr, embedded, omitEmpty bool) (st *sinfo) {
	_ = "STUB: not implemented"
	return nil
}

// reuse previously built

func buildFields(rt reflect.Type, u byte, embedded, omitEmpty bool) (fa []*finfo) {
	_ = "STUB: not implemented"
	return nil
}

func buildTagFields(rt reflect.Type, out, pretty, embedded, omitEmpty bool) (fa []*finfo) {
	_ = "STUB: not implemented"
	return nil
}

func buildExactFields(rt reflect.Type, out, pretty, embedded, omitEmpty bool) (fa []*finfo) {
	_ = "STUB: not implemented"
	return nil
}

func buildLowFields(rt reflect.Type, out, pretty, embedded, omitEmpty bool) (fa []*finfo) {
	_ = "STUB: not implemented"
	return nil
}
