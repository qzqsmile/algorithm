package divideConquer_Tree


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type node struct {
	p *TreeNode
	visited bool
}

func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	stack := []*node{}
	n := node{root, false}
	stack = append(stack, &n)

	for;len(stack)>0;{
		n1 := stack[len(stack)-1]
		stack = stack[0:len(stack)-1]
		if n1.p == nil {
			continue
		}
		if n1.visited == true{
			res = append(res, n1.p.Val)
		}else{
			stack = append(stack, &node{n1.p.Right, false})
			stack = append(stack, &node{n1.p.Left, false})
			n1.visited = true
			stack = append(stack, n1)

		}
	}
	res = append(res)
	return res
}