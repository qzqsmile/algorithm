package divideConquer_Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func delNodes(root *TreeNode, to_delete []int) []*TreeNode {
	mp := make(map[int]bool)
	for _, v := range to_delete {
		mp[v] = true
	}
	res := []*TreeNode{}
	if root != nil && !mp[root.Val] {
		res = append(res, root)
	}
	helper(root, mp, &res)
	return res
}

func helper(root *TreeNode, mp map[int]bool, res *[]*TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root, left, right := root, root.Left, root.Right
	if mp[root.Val] {
		if left != nil && !mp[left.Val] {
			*res = append(*res, left)
		}
		if right != nil && !mp[right.Val] {
			*res = append(*res, right)
		}
		helper(root.Left, mp, res)
		helper(root.Right, mp, res)
		return nil
	} else {
		root.Left = helper(root.Left, mp, res)
		root.Right = helper(root.Right, mp, res)
		return root
	}
}
