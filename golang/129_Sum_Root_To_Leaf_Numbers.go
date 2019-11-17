package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	return helper(root, 0 )
}

func helper(root *TreeNode, sum int) int{
	if root == nil{
		return sum
	}
	left := helper(root.Left, sum*10+root.Val)
	right := helper(root.Right, sum*10+root.Val)
	if root.Left ==  nil{
		return right
	}
	if root.Right == nil{
		return left
	}
	return left+right
}