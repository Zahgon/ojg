// Copyright (c) 2020, Peter Ohler, All rights reserved.

package alt

// String converts the value provided to a string. If conversion is not
// possible such as if the provided value is an array then the first option
// default value is returned or if not provided and empty string is
// returned. If the type is not a string or gen.String and there is a second
// optional default then that second default value is returned. This approach
// keeps the return as a single value and gives the caller the choice of how
// to indicate a bad value.
func String(v any, defaults ...string) (s string) { _ = "STUB: not implemented"; return "" }
