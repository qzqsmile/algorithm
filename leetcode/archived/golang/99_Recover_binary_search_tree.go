package main

type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var res []*TreeNode
	inorderTravser(root, &res)
	var pre, n *TreeNode
	pre, n = nil, nil
	for i := 0; i < len(res)-1; i++ {
		if res[i].Val > res[i+1].Val {
			if pre == nil {
				pre, n = res[i], res[i+1]
			} else {
				n = res[i+1]
			}
		}
	}
	pre.Val, n.Val = n.Val, pre.Val

}

func inorderTravser(cur *TreeNode, res *[]*TreeNode) {
	if cur == nil {
		return
	}
	inorderTravser(cur.Left, res)
	*res = append(*res, cur)
	inorderTravser(cur.Right, res)
}

func main() {

}
