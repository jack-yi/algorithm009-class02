/*
 * @lc app=leetcode.cn id=111 lang=golang
 *
 * [111] 二叉树的最小深度
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return min(minDepth(root.Left), minDepth(root.Right)) + 1
}

func min(x, y int) int {
	if x == 0 && y == 0 {
		return 0
	}
	if x == 0 {
		return y
	}
	if y == 0 {
		return x
	}
	if x <= y {
		return x
	}
	return y
}

// @lc code=end
