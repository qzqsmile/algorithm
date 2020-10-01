package main

func isBalanced (root *TreeNode) bool {
	// write your code here
	l, _ := isBalancedHelper(root)
	return l
}

func isBalancedHelper(root *TreeNode) (bool, int){
	if root == nil{
		return true, 0
	}
	if root.Left == nil && root.Right == nil{
		return true, 1
	}
	lb, ld := isBalancedHelper(root.Left)
	rb, rd := isBalancedHelper(root.Right)

	if lb == true && rb == true && abs(ld-rd) <= 1{
		return true, max(ld, rd)+1
	}
	return false, max(ld, rd)+1
}

func max(a int, b int) int{
	if a > b{
		return a
	}
	return b
}


func abs(a int) int{
	if a < 0{
		return -a
	}
	return a
}
