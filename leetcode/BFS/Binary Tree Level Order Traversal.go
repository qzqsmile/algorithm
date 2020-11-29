package BFS

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int{
	res := [][]int{}
	if root == nil{
		return res
	}
	q := []*TreeNode{}
	q = append(q, root)
	for;len(q) > 0;{
		nq := []*TreeNode{}
		l := []int{}
		for;len(q)>0;{
			n := q[0]
			q = q[1: len(q)]
			l = append(l, n.Val)
			if n.Left != nil{
				nq = append(nq, n.Left)
			}
			if n.Right != nil{
				nq = append(nq, n.Right)
			}
		}
		res = append(res, l)
		q = nq
	}
	return res
}
