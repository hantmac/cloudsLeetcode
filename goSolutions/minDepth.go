package main

import "leetcode/tree"

/*

很多人写出的代码都不符合 1,2 这个测试用例，是因为没搞清楚题意

题目中说明:叶子节点是指没有子节点的节点，这句话的意思是 1 不是叶子节点

题目问的是到叶子节点的最短距离，所以所有返回结果为 1 当然不是这个结果

另外这道题的关键是搞清楚递归结束条件

叶子节点的定义是左孩子和右孩子都为 null 时叫做叶子节点
当 root 节点左右孩子都为空时，返回 1
当 root 节点左右孩子有一个为空时，返回不为空的孩子节点的深度
当 root 节点左右孩子都不为空时，返回左右孩子较小深度的节点值
 */
func minDepth(root *tree.TreeNode) int {
	if root == nil {
		return 0
	}
	var l = 0
	var r = 0

	l = minDepth(root.Left)
	r = minDepth(root.Right)

	if root.Left == nil || root.Right == nil {
		return l + r + 1
	}

	if l <= r {
		l += 1
		return l
	}
	if l > r {
		r += 1
		return r
	}
	return 0
}
