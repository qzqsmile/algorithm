/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

 type NodeIndex struct{
    Node *TreeNode
    Index int
}

func widthOfBinaryTree(root *TreeNode) int {
    if root == nil{
        return 0
    }
    q := []*NodeIndex{&NodeIndex{root, 0}}
    maxWidth := math.MinInt32
    for;len(q) > 0;{
        maxWidth = max(maxWidth, q[len(q)-1].Index-q[0].Index+1)
        t := []*NodeIndex{}
        for;len(q) > 0;{
            n := q[0]
            if n.Node.Left != nil{
                t = append(t, &NodeIndex{n.Node.Left, 2*n.Index+0})
            }
            if n.Node.Right != nil{
                t = append(t, &NodeIndex{n.Node.Right, 2*n.Index+1})
            }
            q = q[1:]
        }
        q = t
    }
    return maxWidth
}


func max(a int, b int) int{
    if a > b{
        return a
    }
    return b
}