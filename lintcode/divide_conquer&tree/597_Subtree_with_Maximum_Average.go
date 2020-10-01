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
 * @param root: the root of binary tree
 * @return: the root of the maximum average of subtree
 */

import "math"

func findSubtree2 (root *TreeNode) *TreeNode {
	// write your code here
	maxAvg := float64(math.MinInt64)
	var maxAvgT *TreeNode
	findSubtree2Helper(&maxAvg, &maxAvgT, root)
	return maxAvgT
}

func findSubtree2Helper(maxAvg *float64,  maxAvgT **TreeNode, root *TreeNode) (int, int){
	if root == nil{
		return 0, 0
	}
	if root.Left == nil && root.Right == nil{
		if float64(root.Val) > *maxAvg{
			*maxAvg = float64(root.Val)
			*maxAvgT = root
		}
		return root.Val, 1
	}

	ls, la := findSubtree2Helper(maxAvg, maxAvgT, root.Left)
	rs, ra := findSubtree2Helper(maxAvg, maxAvgT, root.Right)

	s := ls + rs + root.Val
	n := la + ra + 1
	if float64(s) / float64(n) > *maxAvg{
		*maxAvg = float64(s) / float64(n)
		*maxAvgT = root
	}
	return s, n
}



