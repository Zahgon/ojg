// Copyright (c) 2021, Peter Ohler, All rights reserved.

package sen

// AddMongoFuncs adds TokenFuncs for the common mongo Javascript functions
// that appear in the output from mongosh for some types. They functions
// included are:
//
//	ISODate(arg) returns time.Time when given either a RFC3339 string or milliseconds
//	ObjectId(arg) returns the arg as a string
//	NumberInt(arg)  returns the string argument as an int64 or if too large the original string
//	NumberLong(arg)  returns the string argument as an int64 or if too large the original string
//	NumberDecimal(arg)  returns the string argument as a float64 or if too large the original string
func (p *Parser) AddMongoFuncs() { _ = "STUB: not implemented"; return }

func isoDate(args ...any) (t any) { _ = "STUB: not implemented"; return *new(any) }

func objectID(args ...any) (v any) { _ = "STUB: not implemented"; return *new(any) }

func numberInt64(args ...any) (v any) { _ = "STUB: not implemented"; return *new(any) }

func numberDecimal(args ...any) (v any) { _ = "STUB: not implemented"; return *new(any) }
