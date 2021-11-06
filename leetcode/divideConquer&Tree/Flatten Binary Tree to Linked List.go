/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func flatten(root *TreeNode)  {
    flattenSubTree(root)
}


func flattenSubTree(root *TreeNode) (*TreeNode, *TreeNode){
    if root == nil{
        return nil, nil
    }
    if root.Left == nil && root.Right == nil{
        return root, root
    }
    lh, lt := flattenSubTree(root.Left)
    rh, rt := flattenSubTree(root.Right)
    root.Left = nil
    
    if lh == nil{
        root.Right = rh 
        return root, rt
    }
    if rh == nil{
        root.Right = lh 
        return root, lt
    }
    root.Right = lh 
    lt.Right = rh
    return root, rt
}