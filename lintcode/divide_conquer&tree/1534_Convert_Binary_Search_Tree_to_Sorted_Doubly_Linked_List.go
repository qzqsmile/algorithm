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
 * @param root: root of a tree
 * @return: head node of a doubly linked list
 */

type TreeNode struct {
	Val int
 	Left *TreeNode
	Right *TreeNode
 }

func treeToDoublyList (root *TreeNode) *TreeNode {
	// Write your code here.
	head, tail := helper1(root)
	head.Left = tail
	tail.Right = head
	return head
}

func helper1(root *TreeNode)(*TreeNode, *TreeNode){
	if root == nil{
		return nil, nil
	}
	if root.Left == nil && root.Right == nil{
		return root, root
	}
	lh, lt := helper1(root.Left)
	rh, rt := helper1(root.Right)
	lr, rr := lh, rt

	if lt != nil{
		lt.Right = root
		root.Left = lt
	}else{
		lr = root
	}

	if rh != nil{
		root.Right = rh
		rh.Left = root
	}else{
		rr = root
	}

	return lr, rr
}

func main(){
	r := TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1, Left: nil, Right: nil},
		Right: &TreeNode{Val: 3, Left: nil, Right: nil}}, Right: &TreeNode{Val:5, Left: nil, Right: nil}}
	treeToDoublyList(&r)
}