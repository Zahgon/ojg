// Copyright (c) 2020, Peter Ohler, All rights reserved.

package alt

// Float convert the value provided to a float64. If conversion is not
// possible such as if the provided value is an array then the first option
// default value is returned or if not provided 0.0 is returned. If the type
// is not one of the float types and there is a second optional default then
// that second default value is returned. This approach keeps the return as a
// single value and gives the caller the choice of how to indicate a bad
// value.
func Float(v any, defaults ...float64) (f float64) { _ = "STUB: not implemented"; return 0 }
