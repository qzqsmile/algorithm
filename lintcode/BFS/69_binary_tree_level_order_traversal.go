package main

import "fmt"

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
/**
 * @param root: A Tree
 * @return: Level order a list of lists of integer
 */
func levelOrder (root *TreeNode) [][]int {
	// write your code here
	res :=  make([][]int, 0)
	if root == nil{
		return res
	}
	var q1, q2 []*TreeNode

	q1 = append(q1, root)
	for;len(q1) > 0;{
		r := make([]int, 0)
		for i := 0; i < len(q1); i++{
			r = append(r, q1[i].Val)
			if q1[i].Left != nil{
				q2 = append(q2, q1[i].Left)
			}
			if q1[i].Right != nil{
				q2 = append(q2, q1[i].Right)
			}
		}
		q1 = append([]*TreeNode(nil), q2...)
		q2 = q2[:0]
	}
	return res
}


func main(){
	r := TreeNode{1,nil,nil}
	n2 := TreeNode{2,nil,nil}
	n3 := TreeNode{3,nil,nil}
	n4 := TreeNode{4,nil,nil}
	n5 := TreeNode{5,nil,nil}
	r.Left = &n2
	r.Right = &n3
	n2.Left = &n4
	n2.Right = &n5
	res := levelOrder(&r)
	fmt.Printf("%v", res)
}