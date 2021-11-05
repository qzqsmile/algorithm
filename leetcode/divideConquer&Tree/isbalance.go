/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func isBalanced(root *TreeNode) bool {
    _, res := helper(root)    
    return res
}

func helper(root *TreeNode) (int, bool){
    if root == nil{
        return 0, true
    }
    if root.Left == nil && root.Right == nil{
        return 1, true
    }
    l, lb := helper(root.Left)
    r, rb := helper(root.Right)
    if lb && rb && abs(l-r)<=1{
        return max(l, r)+1, true
    }
    return max(l, r)+1, false
}

func abs(a int) int{
    if a > 0{
        return a 
    }
    return -a
}

func max(a int, b int) int{
    if a > b{
        return a
    }
    return b
}
