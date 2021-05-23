/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

package divideConquer_Tree


type Node struct {
	p   *TreeNode
	existed bool
}

type BSTIterator struct {
	stk []*Node
}

func Constructor(root *TreeNode) BSTIterator {
	var d BSTIterator
	stk := []*Node{}
	if root != nil {
		stk = append(stk, &Node{root, false})
	}
	d.stk = stk
	return d
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	for ; len(this.stk) > 0; {
		n := this.stk[len(this.stk)-1]
		this.stk = this.stk[0 : len(this.stk)-1]
		if n.existed {
			return n.p.Val
		} else {
			if n.p.Right != nil {
				this.stk = append(this.stk, &Node{n.p.Right, false})
			}
			this.stk = append(this.stk, &Node{n.p, true})
			if n.p.Left != nil {
				this.stk = append(this.stk, &Node{n.p.Left, false})
			}
		}
	}
	return -1
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	if len(this.stk) > 0 {
		return true
	}
	return false
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

func main() {

}
