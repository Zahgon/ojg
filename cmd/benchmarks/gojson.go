// Copyright (c) 2020, Peter Ohler, All rights reserved.

package main

import (
	"testing"
)

func goParse(b *testing.B) { _ = "STUB: not implemented"; return }

func goUnmarshalPatient(b *testing.B) { _ = "STUB: not implemented"; return }

func goUnmarshalCatalog(b *testing.B) { _ = "STUB: not implemented"; return }

func goDecodeReader(b *testing.B) { _ = "STUB: not implemented"; return }

func goDecode(b *testing.B) { _ = "STUB: not implemented"; return }

func goParseChan(b *testing.B) { _ = "STUB: not implemented"; return }

// The go json package does not have a chan based result handler so
// fake it to set the baseline for others.

func goValidate(b *testing.B) { _ = "STUB: not implemented"; return }

func goMarshalCatalog(b *testing.B) { _ = "STUB: not implemented"; return }

func goMarshalPatient(b *testing.B) { _ = "STUB: not implemented"; return }

func marshalJSON(b *testing.B) { _ = "STUB: not implemented"; return }

func marshalJSONIndent(b *testing.B) { _ = "STUB: not implemented"; return }

func jsonEncodeIndent(b *testing.B) { _ = "STUB: not implemented"; return }
