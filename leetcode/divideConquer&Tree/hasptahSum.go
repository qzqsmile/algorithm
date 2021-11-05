/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 //树的遍历要注意返回条件
 // root== nil是递归到最nil节点。 root.left==nil && root.Right == nil
 // 则是倒数的结点
 func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil{
        return false
    }
    if root.Left == nil && root.Right == nil{
        if targetSum == root.Val{
            return true
        }
        return false
    }
    
    l := hasPathSum(root.Left, targetSum-root.Val)
    r := hasPathSum(root.Right, targetSum-root.Val)
    return l || r
}
