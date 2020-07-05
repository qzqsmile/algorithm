package divide_conquer_tree


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
 * @return: Preorder in ArrayList which contains node values.
 */
func preorderTraversal (root *TreeNode) []int {
	// write your code here
	res := []int{}
	if root == nil{
		return res
	}
	res = append(res, root.Val)
	res_l := preorderTraversal(root.Left)
	res_r := preorderTraversal(root.Right)
	res = append(res, res_l...)
	res = append(res, res_r...)
	return res
}

