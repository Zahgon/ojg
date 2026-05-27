// Copyright (c) 2020, Peter Ohler, All rights reserved.

package alt

import (
	"time"
)

// Time convert the value provided to a time.Time. If conversion is not
// possible such as if the provided value is an array then the first option
// default value is returned or if not provided zero time is returned. If the
// type is not one of the int or uint types and there is a second optional
// default then that second default value is returned. This approach keeps the
// return as a single value and gives the caller the choice of how to indicate
// a bad value.
func Time(v any, defaults ...time.Time) (t time.Time) {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

// Only good to minutes.

// Only good to microseconds, not nanoseconds.

// Only good to useconds, not nanoseconds.
