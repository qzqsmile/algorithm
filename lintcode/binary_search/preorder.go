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

type Pair struct{
	r *TreeNode
	v bool
}

func preorderTraversal(root *TreeNode) []int {
	var path[]int
	if root == nil{
		return path
	}
	var s []Pair
	s = append(s, Pair{root, false})
	for; len(s) > 0; {
		last := s[len(s)-1]
		root := last.r
		visited := last.v
		s = s[0:len(s)-1]
		if root == nil{
			continue
		}
		if visited{
			path = append(path, root.Val)
		}else{
			s = append(s, Pair{root.Right, false})
			s = append(s, Pair{root.Left, false})
			s = append(s, Pair{root, true})
		}
	}
	return path
}

