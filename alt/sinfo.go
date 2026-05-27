// Copyright (c) 2021, Peter Ohler, All rights reserved.

package alt

import (
	"reflect"
	"sync"

	"github.com/ohler55/ojg"
)

const (
	maskByTag  = byte(0x01)
	maskExact  = byte(0x02) // exact key vs lowwer case first letter
	maskNested = byte(0x04)
	maskSet    = byte(0x08)
)

// sinfo holds reflect information about a struct.
type sinfo struct {
	rt     reflect.Type
	fields [8][]*finfo
}

var (
	structMut sync.Mutex
	// Keyed by the pointer to the type.
	structMap      = map[uintptr]*sinfo{}
	structEmptyMap = map[uintptr]*sinfo{}
)

func (si *sinfo) getFields(o *ojg.Options) []*finfo { _ = "STUB: not implemented"; return nil }

// getSinfo gets the struct information for the provided value. This is use
// internally and is not expected to be used externally.
func getSinfo(v any, omitEmpty bool) (st *sinfo) { _ = "STUB: not implemented"; return nil }

func buildStruct(rt reflect.Type, x uintptr, omitEmpty bool) (st *sinfo) {
	_ = "STUB: not implemented"
	return nil
}

// reuse previously built

func buildFields(rt reflect.Type, u byte, omitEmpty bool) (fa []*finfo) {
	_ = "STUB: not implemented"
	return nil
}

func buildTagFields(rt reflect.Type, nested, omitEmpty bool) (fa []*finfo) {
	_ = "STUB: not implemented"
	return nil
}

func buildExactFields(rt reflect.Type, nested, omitEmpty bool) (fa []*finfo) {
	_ = "STUB: not implemented"
	return nil
}

func buildLowFields(rt reflect.Type, nested, omitEmpty bool) (fa []*finfo) {
	_ = "STUB: not implemented"
	return nil
}
