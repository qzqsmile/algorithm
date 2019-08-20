func maxPathSum(root *TreeNode) int {
	int_max := int(^uint(0) >> 1)
	min := ^int_max
	dfs(root, &min)

	return min
}

func dfs(root *TreeNode, z *int) int {
	if root == nil{
		return 0
	}
	left := Max(0, dfs(root.Left, z))
	right := Max(0, dfs(root.Right, z))
	*z = Max(*z, left+right+root.Val)
	return Max(left, right) + root.Val
}

func Max(x, y int) int{
	if x < y {
		return y
	}
	return x
}