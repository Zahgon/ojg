// Copyright (c) 2020, Peter Ohler, All rights reserved.

package alt

// Int convert the value provided to an int64. If conversion is not possible
// such as if the provided value is an array then the first option default
// value is returned or if not provided 0 is returned. If the type is not one
// of the int or uint types and there is a second optional default then that
// second default value is returned. This approach keeps the return as a
// single value and gives the caller the choice of how to indicate a bad
// value.
func Int(v any, defaults ...int64) (i int64) { _ = "STUB: not implemented"; return 0 }
