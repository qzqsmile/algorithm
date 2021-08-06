package main

import "fmt"

type TreeNode1 struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	var res [][] int
	if root == nil{
		return res
	}
	var q1[] *TreeNode

	q1 = append(q1, root)
	for len(q1) > 0{
		var tmp [] int
		var q2 [] *TreeNode
		for len(q1) > 0{
			t := q1[0]
			tmp = append(tmp, t.Val)
			if t.Left != nil{
				q2 = append(q2, t.Left)
			}
			if t.Right != nil{
				q2 = append(q2, t.Right)
			}
			q1 = q1[1:]
		}
		res = append(res, tmp)
		q1 = q2
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

func main()  {
	//queue := list.New()
	//t := &TreeNode{Val:3}
	//queue.PushBack(t) // Enqueue
	//queue.PushBack(2)
	//
	//for queue.Len() > 0 {
	//	e := queue.Front() // First element
	//	fmt.Print(e.Val)
	//
	//	queue.Remove(e) // Dequeue
	//}
	root := &TreeNode{Val:3}
	root.Left = &TreeNode{Val:9}
	root.Right = &TreeNode{Val:20}
	root.Right.Left = &TreeNode{Val:15}
	root.Right.Right = &TreeNode{Val:7}

	fmt.Println(levelOrderBottom(root))
}
