package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
	var res [][]int
	if root == nil{
		return res
	}
	if root.Left == nil && root.Right == nil{
		if sum == root.Val{
			res = [][]int{[]int{root.Val}}
		}
		return res
	}
	left := pathSum(root.Left, sum-root.Val)
	right := pathSum(root.Right, sum-root.Val)
	for _, v := range left{
		v = append(v[:1], v[0:]...)
		v[0] = root.Val
		res = append(res, v)
	}
	for _, v := range right{
		v = append(v[:1], v[0:]...)
		v[0] = root.Val
		res = append(res, v)
	}
	return res
}

