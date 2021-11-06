/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func kthLargest(root *TreeNode, k int) int {
    res := []int{}
    helper(root, &res)
    return res[k-1]
}

func helper(root *TreeNode, res *[]int){
    if root == nil{
        return 
    }
    helper(root.Right, res)
    *res = append(*res, root.Val)
    helper(root.Left, res)
}