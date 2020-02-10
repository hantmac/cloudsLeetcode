package main

import (
	"leetcode/tree"
)

func sortedArrayToBST2(nums []int, start, end int) *tree.TreeNode {
	if start == end {
		return nil
	}
	var root tree.TreeNode
	var mid = (start + end) / 2
	root = tree.TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST2(nums, start, mid)
	root.Right = sortedArrayToBST2(nums, mid+1, end)
	return &root

}

func sortedArrayToBST(nums []int) *tree.TreeNode {
	return sortedArrayToBST2(nums, 0, len(nums))
}
