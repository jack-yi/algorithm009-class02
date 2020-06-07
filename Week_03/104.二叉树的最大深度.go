/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
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

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	queue := list.New()
	depth := 1
	queue.PushBack(root)
	standByInQueueNode := make([]*TreeNode, 0)
	for queue.Len() != 0 {
		e := queue.Front()
		queue.Remove(e)
		if e != nil {
			tNode := e.Value.(*TreeNode)
			if tNode.Left != nil {
				standByInQueueNode = append(standByInQueueNode, tNode.Left)
			}
			if tNode.Right != nil {
				standByInQueueNode = append(standByInQueueNode, tNode.Right)
			}
		}

		if queue.Len() == 0 && len(standByInQueueNode) > 0 {
			depth++
			for _, v := range standByInQueueNode {
				if v != nil {
					queue.PushBack(v)
				}
			}
			standByInQueueNode = []*TreeNode{}
		}
	}
	return depth 
}

// @lc code=end
