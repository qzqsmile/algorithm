package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	 Val int
	 Left *TreeNode
	 Right *TreeNode
}

type node struct {
	treeNode  *TreeNode
	isVisited bool
}

func (this *BSTIterator) inorderTraversal() int {
	for len(this.stack) > 0 {
		t := this.stack[len(this.stack)-1]
		this.stack = this.stack[0 : len(this.stack)-1]
		if t.isVisited {
			return t.treeNode.Val
		} else {
			if t.treeNode.Right != nil {
				this.stack = append(this.stack, node{t.treeNode.Right, false})
			}
			this.stack = append(this.stack, node{t.treeNode, true})
			if t.treeNode.Left != nil {
				this.stack = append(this.stack, node{t.treeNode.Left, false})
			}
		}
	}
	return 0
}

type BSTIterator struct {
	stack []node
}

func Constructor(root *TreeNode) BSTIterator {
	n := BSTIterator{}
	if root != nil {
		n.stack = append(n.stack, node{root, false})
	}
	return n
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	return this.inorderTraversal()
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	if len(this.stack) > 0 {
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
