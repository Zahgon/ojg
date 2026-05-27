// Copyright (c) 2020, Peter Ohler, All rights reserved.

package alt

// Bool convert the value provided to a bool. If conversion is not possible
// such as if the provided value is an array then the first option default
// value is returned or if not provided false is returned. If the type is not
// a bool nor a gen.Bool and there is a second optional default then that
// second default value is returned. This approach keeps the return as a
// single value and gives the caller the choice of how to indicate a bad
// value.
func Bool(v any, defaults ...bool) (b bool) { _ = "STUB: not implemented"; return false }
