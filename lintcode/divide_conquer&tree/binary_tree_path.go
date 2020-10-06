package main

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths (root *TreeNode) []string {
	// write your code here
	res := []string{}
	helper(&res, root)
	return res
}

func helper(res *[]string, root *TreeNode){
	if root == nil{
		return
	}

	if root.Left == nil && root.Right == nil{
		*res = append(*res, strconv.Itoa(root.Val))
		return
	}

	l, r := []string{}, []string{}
	helper(&l, root.Left)
	helper(&r, root.Right)

	for i, _ := range l{
		*res = append(*res, strconv.Itoa(root.Val)+"->"+l[i])
	}
	for i, _ := range r{
		*res = append(*res, strconv.Itoa(root.Val)+"->"+r[i])
	}
}

