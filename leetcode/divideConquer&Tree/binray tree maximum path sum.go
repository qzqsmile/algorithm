/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 //记忆数组大法好
 func maxPathSum(root *TreeNode) int {
    mp := make(map[*TreeNode]int)
    pathmp := make(map[*TreeNode]int)
    return helper(root, mp, pathmp)
}

func helper(root *TreeNode, mp map[*TreeNode]int, pathmp map[*TreeNode]int) int{
    if _, ok := mp[root]; ok{
        return mp[root]
    }
    if root == nil{
        return -100000
    }
    l := helper(root.Left, mp, pathmp)
    r := helper(root.Right, mp, pathmp)
    l1 := maxPath(root.Left, pathmp)
    r1 := maxPath(root.Right, pathmp)

    res := max(max(l, r),max(max(root.Val,max(l1+root.Val, r1+root.Val)), l1+r1+root.Val))
    mp[root] = res
    return res
}

func maxPath(root *TreeNode, pathmp map[*TreeNode]int) int{
    if _, ok := pathmp[root]; ok{
        return pathmp[root]
    }
    if root == nil {
        return 0
    }
    if root.Left == nil && root.Right == nil{
        return root.Val
    }
    l := maxPath(root.Left, pathmp)
    r := maxPath(root.Right, pathmp)
    res := max(max(l+root.Val, r+root.Val), root.Val)
    pathmp[root] = res
    return res
}

func max(a int, b int) int{
    if a > b{
        return a
    }
    return b
}