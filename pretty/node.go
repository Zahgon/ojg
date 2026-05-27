// Copyright (c) 2021, Peter Ohler, All rights reserved.

package pretty

const (
	strNode   = 's'
	numNode   = 'n'
	arrayNode = 'a'
	mapNode   = 'm'
)

type node struct {
	key     []byte
	members []*node
	buf     []byte
	size    int
	depth   int
	kind    byte
	skip    bool
}

type table struct {
	key     any // string or int
	size    int
	columns []*table
}

func (n *node) subKind() (kind byte) { _ = "STUB: not implemented"; return 0 }

func (n *node) genTables(lazy bool) *table { _ = "STUB: not implemented"; return nil }

func (n *node) updateArrayTable(t *table, lazy bool) { _ = "STUB: not implemented"; return }

func (n *node) updateMapTable(t *table, lazy bool) { _ = "STUB: not implemented"; return }
