/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func diameterOfBinaryTree(root *TreeNode) int {
    longest := math.MinInt32
    helper(root, &longest)
    return longest-1
}

func helper(root *TreeNode, longest *int) int{
    if root == nil{
        return 0
    }
    l := helper(root.Left, longest)
    r := helper(root.Right, longest)
    *longest = max(*longest, l+r+1)
    
    return max(l, r)+1
}

func max(a int, b int) int{
    if a > b{
        return a
    }
    return b
}