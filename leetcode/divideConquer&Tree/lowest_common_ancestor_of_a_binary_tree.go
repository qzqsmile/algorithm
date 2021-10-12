package divideConquer_Tree

/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
// func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
// 	r, _, _ := helper(root, p, q)
// 	return r
// }

// func helper(root, p, q *TreeNode)(*TreeNode, bool, bool){
// 	if root == nil{
// 		return root, false, false
// 	}
// 	rl, lp, lq := helper(root.Left, p, q)
// 	rr, rp, rq := helper(root.Right, p, q)
// 	if rl != nil{
// 		return rl, true, true
// 	}
// 	if rr != nil{
// 		return rr, true, true
// 	}
// 	l := lp || rp || (root.Val == p.Val)
// 	r := lq || rq || (root.Val == q.Val)
// 	var n *TreeNode
// 	if l && r{
// 		n = root
// 	}
// 	return n, l, r
// }


func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil || root == p || root == q{
        return root
    }
    left := lowestCommonAncestor(root.Left, p, q)
    right := lowestCommonAncestor(root.Right, p, q)
    if left != nil && right != nil{
        return root
    }
    if left != nil{
        return left
    }
    if right != nil{
        return right
    }
    return nil
}