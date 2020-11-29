package divideConquer_Tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(s *TreeNode, t *TreeNode) bool {
	if s == nil || t == nil{
		if t == nil{
			return true
		}else{
			return false
		}
	}
	if s.Val == t.Val && isSame(s, t){
		return true
	}else{
		l := isSubtree(s.Left, t)
		r := isSubtree(s.Right, t)
		return l || r
	}
}

func isSame(s *TreeNode, t *TreeNode) bool{
	if s == nil || t == nil{
		if t == nil && s == nil{
			return true
		}else{
			return false
		}
	}
	if s.Val == t.Val{
		l := isSame(s.Left, t.Left)
		r := isSame(s.Right, t.Right)
		return l && r
	}else{
		return false
	}
}
