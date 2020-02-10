package main

import "leetcode/tree"


//实际上就是求最大深度，在每层深度的时候判断下l与r的差值
func isBalanced(root *tree.TreeNode) bool {
	var ans = true
	_,f := depth(root, ans)
	return f
}

func depth(root *tree.TreeNode, f bool) (int,bool) {
	if root == nil {
		return 0,f
	}
	var l = 0
	var r = 0
	l,f = depth(root.Left, f)
	r,f = depth(root.Right, f)
	if l-r > 1 || r-l > 1 {
		f = false
	}
	if l >= r {
		l += 1
		return l,f
	}
	if l < r {
		r += 1
		return r,f
	}
	return 0,f
}
