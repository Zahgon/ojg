// Copyright (c) 2026, Peter Ohler, All rights reserved.

package alt

import (
	"hash/crc64"
)

var emcaTable = crc64.MakeTable(crc64.ECMA)

// Checksum of the provided data using a custom encoding and checksum
// routine. The functions is most efficient with simple data.
func Checksum(v any) uint64 { _ = "STUB: not implemented"; return 0 }

func checksumAppend(b []byte, v any) []byte { _ = "STUB: not implemented"; return nil }

func appendUint64(b []byte, v uint64) []byte { _ = "STUB: not implemented"; return nil }
