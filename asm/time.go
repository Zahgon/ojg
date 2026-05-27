// Copyright (c) 2021, Peter Ohler, All rights reserved.

package asm

func init() {
	Define(&Fn{
		Name: "time?",
		Eval: timeCheck,
		Desc: `Returns true if the single required argumement is a time
otherwise false is returned.`,
	})
	Define(&Fn{
		Name: "time",
		Eval: timeConv,
		Desc: `Converts the first argument to a time if possible otherwise
an error is raised. The first argument can be a integer, float,
or string and are converted as follows:
  integer < 10^10:  time in seconds since 1970-01-01 UTC
  integer >= 10^10: time in nanoseconds 1970-01-01 UTC
  decimal (float):  time in seconds 1970-01-01 UTC
  string:           assumed to be formated as RFC3339 unless a
                    format argument is provided`,
	})
}

func timeCheck(root map[string]any, at any, args ...any) any {
	_ = "STUB: not implemented"
	return *new(any)
}

func timeConv(root map[string]any, at any, args ...any) (t any) {
	_ = "STUB: not implemented"
	return *new(any)
}
