package divideConquer_Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//这题唯一的confuse的地方在于理解左叶子，其实就需要满足两点
//1 是叶子节点，即度为0 2.在一颗左子树上。 理解这两点这题就很easy
func sumOfLeftLeaves(root *TreeNode) int {
	return helper(root, false)
}

func helper(root *TreeNode, isLeft bool) int{
	if root == nil{
		return 0
	}
	if root.Left == nil &&
		root.Right == nil && isLeft{
		return root.Val
	}
	l := helper(root.Left, true)
	r := helper(root.Right, false)
	return l + r
}

