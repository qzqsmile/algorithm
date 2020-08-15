package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
 }
/**
 * @param root: A Tree
 * @return: Level order a list of lists of integer
 */
func levelOrder (root *TreeNode) [][]int {
	// write your code here
	res :=  make([][]int, 0)
	if root == nil{
		return res
	}
	var q1, q2 []*TreeNode

	q1 = append(q1, root)
	for;len(q1) > 0;{
		r := make([]int, 0)
		for i := 0; i < len(q1); i++{
			r = append(r, q1[i].Val)
			if q1[i].Left != nil{
				q2 = append(q2, q1[i].Left)
			}
			if q1[i].Right != nil{
				q2 = append(q2, q1[i].Right)
			}
		}
		q1 = append([]*TreeNode(nil), q2...)
		q2 = q2[:0]
	}
	return res
}

