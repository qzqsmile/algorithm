package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil{
		return nil
	}
	if root.Val == key{
		l := root.Left
		r := root.Right
		if r == nil{
			return l
		}else if r.Left == nil{
			r.Left = l
			return r
		}else{
			nr := r.Left
			for nr.Left != nil{
				r = nr
				nr = nr.Left
			}
			r.Left = nr.Right
			nr.Left = root.Left
			nr.Right = root.Right
			return nr
		}
	}else if key < root.Val{
		root.Left = deleteNode(root.Left, key)
	}else{
		root.Right = deleteNode(root.Right, key)
	}
	return root
}
