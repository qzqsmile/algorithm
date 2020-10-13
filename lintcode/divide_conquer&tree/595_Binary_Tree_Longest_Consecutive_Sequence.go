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
 * @return: the length of the longest consecutive sequence path
 */
func longestConsecutive (root *TreeNode) int {
	// write your code here
	maxCount := 0
	helper(root, &maxCount)
	return maxCount
}

func helper(root *TreeNode, maxCount *int) (int){
	if root == nil{
		return 0
	}
	if root.Left == nil && root.Right == nil{
		if 1 > *maxCount{
			*maxCount = 1
		}
		return 1
	}

	nl := helper(root.Left, maxCount)
	tr := helper(root.Right, maxCount)
	c := 1
	if root.Left != nil && root.Val+1 == root.Left.Val{
		c = max(nl+1, c)
	}
	if root.Right != nil && root.Val+1 == root.Right.Val{
		c = max(tr+1, c)
	}
	*maxCount = max(*maxCount, c)
	return c
}

func max(a int, b int) int{
	if a > b{
		return a
	}
	return b
}