package main
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
 * @param root: A Tree
 * @return: Inorder in ArrayList which contains node values.
 */
func inorderTraversal (root *TreeNode) []int {
	// write your code here
	res := []int{}
	if root == nil{
		return res
	}
	res_l := inorderTraversal(root.Left)
	res = append(res, res_l...)
	res = append(res, root.Val)
	res_r := inorderTraversal(root.Right)
	res = append(res, res_r...)
	return res
}
