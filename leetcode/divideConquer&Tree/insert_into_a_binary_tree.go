package divideConquer_Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil{
		return &TreeNode{val, nil, nil}
	}
	if val > root.Val{
		if root.Right == nil{
			root.Right = &TreeNode{val, nil , nil}
		}else{
			insertIntoBST(root.Right, val)
		}
	}else{
		if root.Left == nil{
			root.Left = &TreeNode{val, nil, nil}
		}else{
			insertIntoBST(root.Left, val)
		}
	}
	return root
}