/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 func rightSideView(root *TreeNode) []int {
    res := []int{}
    if root == nil{
        return res
    }
    q := []*TreeNode{}
    q = append(q, root)

    for;len(q) > 0;{
        res = append(res, q[len(q)-1].Val)
        nq := []*TreeNode{}
        for;len(q) > 0;{
            n := q[0]
            if n.Left != nil{
                nq = append(nq, n.Left)
            }
            if n.Right != nil{
                nq = append(nq, n.Right)
            }
            q = q[1:len(q)]
        }
        q = nq
    }

    return res
}