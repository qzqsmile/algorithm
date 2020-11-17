package divideConquer_Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func flatten(root *TreeNode)  {
	if root == nil{
		return
	}
	helper(root)
}

func helper(root *TreeNode) (*TreeNode, *TreeNode){
	if root.Left == nil && root.Right == nil{
		return root, root
	}else if root.Left == nil{
		rh, rt := helper(root.Right)
		root.Right = rh
		root.Left = nil
		return root, rt
	}else if root.Right == nil{
		lh, lt := helper(root.Left)
		root.Right = lh
		root.Left = nil
		return root, lt
	}else{
		rh, rt := helper(root.Right)
		lh, lt := helper(root.Left)
		root.Right = lh
		root.Left = nil
		lt.Right = rh
		return root, rt
	}
}