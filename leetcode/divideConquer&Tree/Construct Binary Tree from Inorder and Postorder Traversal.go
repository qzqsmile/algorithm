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

我注意到大部分的章节都是先用一些例子起篇，这样可以帮助读者先建立一个直观的理解。但不知道为什么5.3 context却先从实现原理入手，然后再给出例子。这样对于我这种对context都不了解的人会看的云里雾里。作者这里是有什么特殊的考量么？

