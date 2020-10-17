package main

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

type NestedIterator struct {
	stk []*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	n := NestedIterator{stk: []*NestedInteger{}}
	for i := len(nestedList) - 1; i >= 0; i-- {
		n.stk = append(n.stk, nestedList[i])
	}
	return &n
}

func (this *NestedIterator) Next() int {
	top := this.stk[len(this.stk)-1]
	this.stk = this.stk[0 : len(this.stk)-1]
	return top.GetInteger()
}

func (this *NestedIterator) HasNext() bool {
	for len(this.stk) > 0 {
		topList := this.stk[len(this.stk)-1].GetList()
		if len(topList) > 0 {
			this.stk = this.stk[0 : len(this.stk)-1]
			for i := len(topList) - 1; i >= 0; i-- {
				this.stk = append(this.stk, topList[i])
			}
		} else {
			top := this.stk[len(this.stk)-1]
			if top.IsInteger() {
				return true
			} else {
				this.stk = this.stk[0 : len(this.stk)-1]
			}
		}
	}
	return false
}
