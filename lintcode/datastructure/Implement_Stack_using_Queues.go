package main


type MyStack struct {
	q1 []int
	q2 []int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	m := MyStack{q1:[]int{}, q2:[]int{}}
	return m
}

/** Push element x onto stack. */

func (this *MyStack) Push(x int)  {
	this.q1 = append(this.q1, x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	for len(this.q1) > 1{
		t := this.q1[0]
		this.q1 = this.q1[1:]
		this.q2 = append(this.q2, t)
	}
	r := this.q1[0]
	this.q1 = this.q1[1:]
	for len(this.q2) >0{
		t := this.q2[0]
		this.q2 = this.q2[1:]
		this.q1 = append(this.q1, t)
	}
	return r
}


/** Get the top element. */
func (this *MyStack) Top() int {
	r := this.Pop()
	this.Push(r)
	return r
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if len(this.q1) > 0{
		return true
	}
	return false
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */