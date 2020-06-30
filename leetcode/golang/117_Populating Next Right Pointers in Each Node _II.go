/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

 func connect(root *Node) *Node {
    r := root
    q := make([] *Node, 0)
    if root != nil{
        q = append(q, root)
    }
    for ;len(q) > 0; {
        q1 := make([] *Node, 0)
        for i := 0; i < len(q); i++{
            if i == len(q)-1{
                q[i].Next = nil
            }else{
                q[i].Next = q[i+1]
            }
            if q[i].Left != nil{
                q1 = append(q1, q[i].Left)
            }
            if q[i].Right != nil{
                q1 = append(q1, q[i].Right)
            }
        }
        q = q1
    }
    return r
}