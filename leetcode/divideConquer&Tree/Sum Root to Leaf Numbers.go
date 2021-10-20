/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	return helper(root, 0)
}

func helper(root *TreeNode, parentSum int) int {
	if root.Left == nil && root.Right == nil {
		return 10*parentSum + root.Val
	}
	l, r := 0, 0
	if root.Left != nil {
		l = helper(root.Left, 10*parentSum+root.Val)
	}
	if root.Right != nil {
		r = helper(root.Right, 10*parentSum+root.Val)
	}

	return l + r
}