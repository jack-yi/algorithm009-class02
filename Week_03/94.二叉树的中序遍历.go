/*
 * @lc app=leetcode.cn id=94 lang=golang
 *
 * [94] 二叉树的中序遍历
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
var r []int

func inorderTraversal(root *TreeNode) []int {
	r = make([]int, 0)
	stack := make([]*TreeNode, 0)

	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		index := len(stack) - 1
		r = append(r, stack[index].Val)
		root = stack[index].Right
		stack = stack[:index]
	}
	return r
}

// @lc code=end
