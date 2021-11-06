/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func invertTree(root *TreeNode) *TreeNode {
    if root == nil{
        return root
    }
    r, l := root.Right, root.Left
    root.Left = invertTree(r)
    root.Right = invertTree(l)
    // root.Left, root.Right = root.Right, root.Left 
    return root
}