/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func pathSum(root *TreeNode, targetSum int) [][]int {
    res := [][]int{}
    tmp := []int{}
    helper(root, targetSum, tmp, &res)
    return res
}

func helper(root *TreeNode, targetSum int, tmp []int, res *[][]int){
    if root == nil{
        return 
    }        
    tmp = append(tmp, root.Val)
    if root.Left == nil && root.Right == nil{
        if targetSum == root.Val && len(tmp) != 0{
            *res = append(*res, append([]int{}, tmp...))
        }
    }
    helper(root.Left, targetSum-root.Val, tmp, res)
    helper(root.Right, targetSum-root.Val, tmp, res)
    tmp = tmp[0:len(tmp)-1]
}