package divideConquer_Tree



/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//solution1 travers+每个根节点判断的思想。很直观，但是时间复杂度很多O(n^2)。没想到
//这个算法居然能Pass
func pathSum(root *TreeNode, sum int) int {
	s := 0
	traverse(root, &s, sum)
	return s
}
func traverse(root *TreeNode, s *int, sum int){
	if root == nil{
		return
	}
	r := helper(root, sum)
	*s += r
	traverse(root.Left, s, sum)
	traverse(root.Right, s, sum)
}
func helper(root *TreeNode, sum int) int{
	if root == nil{
		return 0
	}
	res := 0
	if root.Val == sum {
		res = 1
	}
	l := helper(root.Left, sum-root.Val)
	r := helper(root.Right, sum-root.Val)
	return l+r+res
}