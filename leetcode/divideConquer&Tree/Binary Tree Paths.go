package divideConquer_Tree

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	res := []string{}
	if root == nil{
		return res
	}
	if root.Left == nil && root.Right == nil{
		return []string{strconv.Itoa(root.Val)}
	}
	lres := binaryTreePaths(root.Left)
	rres := binaryTreePaths(root.Right)

	for _, v := range lres{
		res = append(res, strconv.Itoa(root.Val)+"->"+v)
	}
	for _, v := range rres{
		res = append(res, strconv.Itoa(root.Val)+"->"+v)
	}
	return res
}
