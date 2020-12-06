package BFS


/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil{
		return res
	}
	q := []*TreeNode{}
	q1 := []*TreeNode{}
	q = append(q, root)
	level := 1
	for;len(q) > 0;{
		q1 = []*TreeNode{}
		tmp := []int{}
		for;len(q) > 0;{
			n := q[0]
			q = q[1:]
			tmp = append(tmp, n.Val)
			if n.Left != nil{
				q1  = append(q1, n.Left)
			}
			if n.Right != nil{
				q1 = append(q1, n.Right)
			}
		}
		q = q1
		if level % 2 == 1{
			res = append(res, tmp)
		}else{
			reverse(tmp)
			res = append(res, tmp)
		}
		level++
	}
	return res
}

func reverse(data []int){
	i, j := 0, len(data)-1
	for;i < j;{
		data[i], data[j] = data[j], data[i]
		i++
		j--
	}
}