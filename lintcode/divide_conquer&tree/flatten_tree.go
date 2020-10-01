package main

func flatten(root *TreeNode)  {
	if root == nil{
		return
	}
	if root.Left == nil && root.Right == nil{
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	l, r := root.Left, root.Right
	if l != nil{
		lr := l
		for;lr.Right != nil; lr = lr.Right{
		}
		root.Right = l
		root.Left = nil
		lr.Right = r
	}
}
