package main

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	_, _, c := lowestCommonAncestorHelper(root, p, q)
	return c
}

func lowestCommonAncestorHelper(root, p, q *TreeNode) (bool, bool, *TreeNode) {
	if root == nil {
		return false, false, root
	}
	lp, lq, lc := lowestCommonAncestorHelper(root.Left, p, q)
	rp, rq, rc := lowestCommonAncestorHelper(root.Right, p, q)
	if lc != nil {
		return true, true, lc
	}
	if rc != nil {
		return true, true, rc
	}
	bp := lp || rp
	bq := lq || rq
	if root == p {
		bp = true
	}
	if root == q {
		bq = true
	}
	if bp && bq {
		return true, true, root
	}
	return bp, bq, nil
}
