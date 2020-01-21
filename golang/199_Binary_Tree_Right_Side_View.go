/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func rightSideView(root *TreeNode) []int {
    q := make([] *TreeNode, 0)
    res := make([]int, 0)
    if root == nil{
        return res;
    }
    q = append(q, root)
    for ;len(q) > 0; {
        q1 := make([] *TreeNode, 0)
        res = append(res, q[0].Val)
        for ;len(q) > 0; {
            tmp := q[0]
            q = q[1:]
            if tmp.Right != nil{
                q1 = append(q1, tmp.Right)
            }
            if tmp.Left != nil{
                q1 = append(q1, tmp.Left)
            }
        }
        q = q1
    }
    return res
}

