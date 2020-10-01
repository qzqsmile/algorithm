package main

import "math"

func findSubtree (root *TreeNode) *TreeNode {
	// write your code here
	if root == nil{
		return nil
	}
	minNum := math.MaxInt64
	var minRoot *TreeNode
	findMinSub(&minNum, &minRoot, root)
	return minRoot
}

func findMinSub(minNum *int, t **TreeNode, root *TreeNode) int{
	if root == nil{
		return 0
	}
	if root.Left == nil && root.Right == nil{
		if root.Val < *minNum{
			*minNum = root.Val
			*t = root
		}
		return root.Val
	}
	l := findMinSub(minNum, t, root.Left)
	r := findMinSub(minNum, t, root.Right)
	s := l + r + root.Val
	if s < *minNum{
		*minNum = s
		*t = root
	}
	return s
}
