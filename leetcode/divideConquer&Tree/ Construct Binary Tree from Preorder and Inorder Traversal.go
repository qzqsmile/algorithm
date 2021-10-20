package divideConquer_Tree
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 //本题要注意访问数组不要越界
 func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0{
        return nil
    }

    root := &TreeNode{preorder[0], nil, nil}
    l := 0
    for i, v := range inorder{
        if v == root.Val{
            l = i 
            break
        }
    }
    if l != 0{
        root.Left = buildTree(preorder[1:1+l], inorder[0:l])
    }
    if 1+l != len(preorder){
        root.Right = buildTree(preorder[1+l:], inorder[l+1:])
    }

    return root
}