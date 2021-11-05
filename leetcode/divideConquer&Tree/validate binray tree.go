/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isValidBST(root *TreeNode) bool {
    return helper(root, math.MinInt64, math.MaxInt64)
}

func helper(root *TreeNode, l int, h int) bool{
    if root == nil{
        return true
    }
    if l < root.Val && root.Val < h{
        l := helper(root.Left, l, root.Val)
        r := helper(root.Right, root.Val, h)
        return l && r
    }
    return false
}