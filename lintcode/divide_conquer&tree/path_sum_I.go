package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 * leetcode 112
 */
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil{
		return false
	}
	if root.Left == nil && root.Right == nil{
		if sum == root.Val{
			return true
		}
		return false
	}
	left := hasPathSum(root.Left, sum-root.Val)
	right := hasPathSum(root.Right, sum-root.Val)
	if left || right{
		return true
	}
	return false
}

