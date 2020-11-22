package divideConquer_Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil{
		return false
	}
	if root.Left == nil && root.Right == nil{
		if root.Val == sum{
			return true
		}
		return false
	}
	l := hasPathSum(root.Left, sum - root.Val)
	r := hasPathSum(root.Right, sum - root.Val)
	return l || r
}
