package divideConquer_Tree


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	mp := make(map[int]bool)
	return bfs(root, k, mp)
}

func bfs(root *TreeNode, k int, mp map[int]bool) bool{
	if root == nil{
		return false
	}
	if _, ok := mp[k-root.Val]; ok{
		return true
	}else{
		mp[root.Val] = true
	}
	l := bfs(root.Left, k, mp)
	r := bfs(root.Right, k, mp)
	return l || r
}

