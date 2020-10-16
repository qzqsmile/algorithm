package main

type MyQueue struct {
	stk1 []int
	stk2 []int
	p int
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
	m := MyQueue{stk1: []int{}, stk2: []int{}, p:0}
	return m
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
	this.stk1 = append(this.stk1, x)
	if len(this.stk1) == 1{
		this.p = x
	}
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	for len(this.stk1) > 0{
		this.stk2 = append(this.stk2, this.stk1[len(this.stk1)-1])
		this.stk1 = this.stk1[0: len(this.stk1)-1]
	}
	r := this.stk2[len(this.stk2)-1]
	this.stk2 = this.stk2[0:len(this.stk2)-1]
	if len(this.stk2) > 0{
		this.p = this.stk2[len(this.stk2)-1]
	}
	for len(this.stk2) > 0 {
		this.stk1 = append(this.stk1, this.stk2[len(this.stk2)-1])
		this.stk2 = this.stk2[0:len(this.stk2)-1]
	}
	return r
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
	return this.p
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if len(this.stk1) > 0{
		return false
	}
	return true
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
