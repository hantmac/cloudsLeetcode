package main

import (
	"leetcode/stack"
	"leetcode/tree"
)

//使用递归实现
func maxDepth(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}

	var l = 0
	var r = 0
	l = maxDepth(root.Left)
	r = maxDepth(root.Right)

	if l > r {
		l += 1
		return l
	}
	if l <= r {
		r += 1
		return r
	}
	return 0
}

func maxDepth2(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}
	var depth = 0
	var l = 0
	var r = 0
	s := stack.New()
	s.PUsh(root)
	depth += 1
	for s.Len() != 0 {
		s.Pop()
		if root.Left != nil {
			s.PUsh(root.Left)
			l += 1
		}
		if root.Right != nil {
			s.PUsh(root.Right)
			r += 1
		}
		if l >= r {
			depth = l
		} else {
			depth = r
		}

	}
	return depth
}
