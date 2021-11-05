/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 package divideConquer_Tree
 
 func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {
        return 0 
    }
    l := diameterOfBinaryTree(root.Left)
    r := diameterOfBinaryTree(root.Right)
    l1 := depth(root.Left)
    r1 := depth(root.Right)
    return max(l1+r1, max(l, r))
}

func depth(root *TreeNode) int{
    if root == nil{
        return 0
    }
    return 1 + max(depth(root.Left), depth(root.Right))
}

func max(a int, b int) int{
    if a > b{
        return a
    }
    return b
}