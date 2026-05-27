// Copyright (c) 2020, Peter Ohler, All rights reserved.

package jp

// Walk data and call the cb callback for each node in the data. The path is
// reused in each call so if the path needs to be save it should be copied.
func Walk(data any, cb func(path Expr, value any), justLeaves ...bool) {
	_ = "STUB: not implemented"
	return
}

func walk(path Expr, data any, cb func(path Expr, value any), justLeaves bool) {
	_ = "STUB: not implemented"
	return
}

// leaf node

// Walk the matching elements in the data and call cb on the matches. The path
// passed to the cb function is the normalized path to the current location
// while the nodes are the chain of elements up to and including the current
// element.
func (x Expr) Walk(data any, cb func(path Expr, nodes []any)) { _ = "STUB: not implemented"; return }
