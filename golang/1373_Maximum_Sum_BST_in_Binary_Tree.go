/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxSumBST(root *TreeNode) int {
	max := 0
	dSum(root, &max)
	return max
}

func dSum(root *TreeNode, max *int) (bool, int, int, int) {
	if root == nil {
		return true, 40000, -40000, 0
	}
	if root.Left == nil && root.Right == nil {
		*max = mymax(*max, root.Val)
		return true, root.Val, root.Val, root.Val
	}
	l_is, l_min, l_max, l_sum := dSum(root.Left, max)
	r_is, r_min, r_max, r_sum := dSum(root.Right, max)
	if l_is && r_is {
		if root.Val > l_max && root.Val < r_min {
			tmp := root.Val + l_sum + r_sum
			*max = mymax(*max, tmp)
			min_num := mymin(l_min, root.Val)
			max_num := mymax(r_max, root.Val)
			return true, min_num, max_num, tmp
		}
	}
	return false, 0, 0, 0
}

func mymin(a int, b int) int {
	if (a < b) {
		return a
	}
	return b
}

func mymax(a int, b int) int {
	if (a > b) {
		return a
	}
	return b
}
