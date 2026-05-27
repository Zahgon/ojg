// Copyright (c) 2024, Peter Ohler, All rights reserved.

package jp

// PathMatch returns true if the provided path would match the target
// expression. The path argument is expected to be a normalized path with only
// elements of Root ($), At (@), Child (string), or Nth (int). A Filter
// fragment in the target expression will match any value in path since it
// requires data from a JSON document to be evaluated. Slice fragments always
// return true as long as the path element is an Nth.
func PathMatch(target, path Expr) bool { _ = "STUB: not implemented"; return false }

// ignore and don't advance path

// Assume a match since there is no data for comparison.
