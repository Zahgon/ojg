// Copyright (c) 2021, Peter Ohler, All rights reserved.

package ojg

import (
	"strconv"
	"time"
)

// 23 for fraction in IEEE 754 which amounts to 7 significant digits. Use base
// 10 so that numbers look correct when displayed in base 10.
const fracMax = 10000000.0

// Converter types are used to convert data element to alternate
// values. Common uses are to match a pattern such as strings representing
// dates to time.Time.
type Converter struct {
	// Int are a slice of functions to match and convert Ints.
	Int []func(val int64) (any, bool)

	// Float are a slice of functions to match and convert Floats.
	Float []func(val float64) (any, bool)

	// String are a slice of functions to match and convert Strings.
	String []func(val string) (any, bool)

	// Map are a slice of functions to match and convert Maps.
	Map []func(val map[string]any) (any, bool)

	// Array are a slice of functions to match and convert Arrays.
	Array []func(val []any) (any, bool)
}

var (
	// TimeRFC3339Converter converts strings matching time.RFC3339Nano,
	// time.RFC3339, or 2006-01-02 to time.Time.
	TimeRFC3339Converter = Converter{
		String: []func(val string) (any, bool){
			func(val string) (any, bool) {
				if 20 <= len(val) && len(val) <= 35 {
					for _, layout := range []string{time.RFC3339Nano, time.RFC3339} {
						if t, err := time.ParseInLocation(layout, val, time.UTC); err == nil {
							return t, true
						}
					}
				} else if len(val) == 10 {
					if t, err := time.ParseInLocation("2006-01-02", val, time.UTC); err == nil {
						return t, true
					}
				}
				return val, false
			},
		},
	}

	// TimeNanoConverter converts large integers, 946684800000000000
	// (2000-01-01) and above to time.Time.
	TimeNanoConverter = Converter{
		Int: []func(val int64) (any, bool){
			func(val int64) (any, bool) {
				if 946684800000000000 <= val { // 2000-01-01
					return time.Unix(0, val), true
				}
				return val, false
			},
		},
	}

	// MongoConverter convert maps with one member when the member key is
	// $numberLong, $date, $numberDecimal, or $oid and the value and the
	// member value is a string. These patterns are found in mongodb JSON
	// exports.
	MongoConverter = Converter{
		Map: []func(val map[string]any) (any, bool){
			func(val map[string]any) (any, bool) {
				if len(val) != 1 {
					return val, false
				}
				for k, v := range val {
					s, ok := v.(string)
					if !ok {
						break
					}
					switch k {
					case "$numberLong":
						if i, err := strconv.ParseInt(s, 10, 64); err == nil {
							return i, true
						}
					case "$date":
						if t, err := time.ParseInLocation("2006-01-02T15:04:05.999Z07:00", s, time.UTC); err == nil {
							return t, true
						}
					case "$numberDecimal":
						if f, err := strconv.ParseFloat(s, 64); err == nil {
							return f, true
						}
					case "$oid":
						return s, true
					}
				}
				return val, false
			},
		},
	}
)

// Convert a value according to the conversion functions of the converter. If
// the value is a map or slice and not converted itself the provided value
// will remain the same but will be modified if any of it's members are
// converted.
func (c *Converter) Convert(v any) any { _ = "STUB: not implemented"; return *new(any) }

func (c *Converter) convert(v any) (any, bool) { _ = "STUB: not implemented"; return *new(any), false }

// This small rounding makes the conversion from 32 bit to 64 bit
// display nicer.

// Convert a value according to the conversion functions provided. If the
// value is a map or slice and not converted itself the provided value will
// remain the same but will be modified if any of it's members are converted.
func Convert(v any, funcs ...any) any { _ = "STUB: not implemented"; return *new(any) }
