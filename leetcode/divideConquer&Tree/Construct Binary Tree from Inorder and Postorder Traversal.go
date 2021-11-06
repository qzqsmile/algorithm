/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func buildTree(inorder []int, postorder []int) *TreeNode {
    if len(inorder) == 0{
        return nil
    }
    rootVal := postorder[len(postorder)-1]
    rootIndex := -1
    root := &TreeNode{rootVal, nil, nil}
    for i := 0; i < len(inorder); i++{
        if inorder[i] == rootVal{
            rootIndex = i
            break
        }
    }
    root.Left = buildTree(inorder[:rootIndex], postorder[:rootIndex])
    root.Right = buildTree(inorder[rootIndex+1:], postorder[rootIndex:len(postorder)-1])
    return root
}