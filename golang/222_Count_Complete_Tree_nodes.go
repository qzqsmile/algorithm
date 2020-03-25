/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
	l := leftDepth(root)
	r := rightDepth(root)
	if l == r {
		return int(math.Pow(2, float64(l))) - 1
	} else {
		return 1 + countNodes(root.Left) + countNodes(root.Right)
	}
}

func leftDepth(root *TreeNode) int {
	d := 0
	for l := root; l != nil; l = l.Left {
		d += 1
	}
	return d
}

func rightDepth(root *TreeNode) int {
	d := 0
	for r := root; r != nil; r = r.Right {
		d += 1
	}
	return d
}
