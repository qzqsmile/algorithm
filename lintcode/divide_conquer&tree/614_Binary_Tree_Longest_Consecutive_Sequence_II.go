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
func longestConsecutive2 (root *TreeNode) int {
	// write your code here
	maxCount := 0
	helper(root, &maxCount)
	return maxCount
}

// asc desc
func helper(root *TreeNode, maxCount *int)(int, int){
	if root == nil{
		return 0, 0
	}
	if root.Left == nil && root.Right == nil{
		if 1 > *maxCount{
			*maxCount = 1
		}
		return 1, 1
	}
	la, ld := helper(root.Left, maxCount)
	ra, rd := helper(root.Right, maxCount)
	ca, cd := 1, 1
	if root.Left != nil  && root.Left.Val+1 == root.Val{
		ca = max(la+1, ca)
	}
	if root.Right != nil && root.Right.Val+1 == root.Val{
		ca = max(ra+1, ca)
	}
	if root.Left != nil && root.Left.Val-1 == root.Val{
		cd = max(ld+1, cd)
	}
	if root.Right != nil && root.Right.Val-1 == root.Val{
		cd = max(rd+1, cd)
	}
	*maxCount = max(*maxCount, ca)
	*maxCount = max(*maxCount, cd)
	if root.Left != nil && root.Right != nil{
		if root.Left.Val+1 == root.Val && root.Right.Val-1 == root.Val{
			*maxCount = max(*maxCount, la+rd+1)
		}
		if root.Left.Val-1 == root.Val && root.Right.Val+1 == root.Val{
			*maxCount = max(*maxCount, ld+ra+1)
		}
	}
	return ca, cd
}

func max(a int, b int) int{
	if a > b{
		return a
	}
	return b
}


