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
 * @return: Postorder in ArrayList which contains node values.
 */
func postorderTraversal (root *TreeNode) []int {
	// write your code here
	res := []int{}
	if root == nil{
		return res
	}
	res_l := postorderTraversal(root.Left)
	res = append(res, res_l...)
	res_r := postorderTraversal(root.Right)
	res = append(res, res_r...)
	res = append(res, root.Val)
	return res
}



