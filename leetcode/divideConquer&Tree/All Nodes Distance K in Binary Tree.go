package divideConquer_Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	mp := make(map[int]*TreeNode)
	dfs(root, mp)
	res := []int{}
	targetKth(target, target, k, 0, &res, mp)
	return res
}

func targetKth(node *TreeNode, from *TreeNode, k int, depth int, res *[]int, mp map[int]*TreeNode) {
	if node == nil || depth > k {
		return
	}
	if k == depth {
		*res = append(*res, node.Val)
		return
	}
	if node.Left != from {
		targetKth(node.Left, node, k, depth+1, res, mp)
	}
	if node.Right != from {
		targetKth(node.Right, node, k, depth+1, res, mp)
	}
	if mp[node.Val] != from {
		targetKth(mp[node.Val], node, k, depth+1, res, mp)
	}
}

func dfs(root *TreeNode, mp map[int]*TreeNode) {
	if root == nil {
		return
	}
	if root.Left != nil {
		mp[root.Left.Val] = root
		dfs(root.Left, mp)
	}
	if root.Right != nil {
		mp[root.Right.Val] = root
		dfs(root.Right, mp)
	}
}
