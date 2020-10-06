package main

import "math"

func isValidBST (root *TreeNode) bool {
	// write your code here
	return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, min int, max int) bool{
	if root == nil{
		return true
	}
	if root.Val > min && root.Val < max{
		l := helper(root.Left, min, root.Val)
		r := helper(root.Right, root.Val, max)
		return l && r
	}else{
		return false
	}
}

