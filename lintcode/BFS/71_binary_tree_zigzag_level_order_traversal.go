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
 * @return: A list of lists of integer include the zigzag level order traversal of its nodes' values.
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	// write your code here
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	var q1, q2 []*TreeNode
	q1 = append(q1, root)
	count := 1
	for ; len(q1) > 0; {
		r := make([]int, 0)
		for i := 0; i < len(q1); i++ {
			r = append(r, q1[i].Val)
			if q1[i].Left != nil {
				q2 = append(q2, q1[i].Left)
			}
			if q1[i].Right != nil {
				q2 = append(q2, q1[i].Right)
			}
		}
		if count%2 == 0 {
			reverse(r)
		}
		res = append(res, r)
		count++
		q1 = append([]*TreeNode{}, q2...)
		q2 = q2[:0]
	}
	return res
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
