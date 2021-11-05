package divideConquer_Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//此题的本质是左右子树的不同遍历顺序是否完全相等
func isSymmetric(root *TreeNode) bool {
	return isequal(root.Left, root.Right)
}

func isequal(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	c1 := isequal(left.Right, right.Left)
	c2 := isequal(left.Left, right.Right)

	return c1 && c2
}
