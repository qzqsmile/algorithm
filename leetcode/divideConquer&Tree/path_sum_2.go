package divideConquer_Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func pathSum(root *TreeNode, sum int) [][]int {
	res := [][]int{}
	if root == nil{
		return res
	}
	if root.Left == nil && root.Right == nil && root.Val == sum{
		return [][]int{[]int{root.Val}}
	}
	l := pathSum(root.Left, sum-root.Val)
	r := pathSum(root.Right, sum-root.Val)
	for _, v := range l{
		n := append([]int{root.Val}, v...)
		res = append(res, n)
	}
	for _, v := range r{
		n := append([]int{root.Val}, v...)
		res = append(res, n)
	}
	return res
}
