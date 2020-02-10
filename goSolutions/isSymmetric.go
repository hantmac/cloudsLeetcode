package main

import "leetcode/tree"

func isSymmetric(root *tree.TreeNode) bool {

	return check(root, root)
}

func check(p *tree.TreeNode, q *tree.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p != nil || q != nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	return check(p.Left, q.Right) && check(p.Right, q.Left)
}
